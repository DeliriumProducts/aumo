import axios from "axios"
import ShopItemList from "../components/ShopItemList"
import * as SecureStore from "expo-secure-store"
import Receipt from "../components/Receipt"
import React from "react"
import {
  ScrollView,
  StyleSheet,
  Text,
  View,
  ActivityIndicator
} from "react-native"
import { BACKEND_URL } from "../config"
import { Avatar, Button, TabView, Tab } from "react-native-ui-kitten"
import { withNavigationFocus } from "react-navigation"

const receiptContent = `Портокалов сок 1L x 1: 1,99лв.

Сладолед ванилия 80гр x 4: 4,20лв.

Спагети 700гр x 1 : 3,30лв.`

const UserScreen = props => {
  const [mode, setMode] = React.useState("rewards")
  const [user, setUser] = React.useState(null)
  const [loading, setLoading] = React.useState(true)

  React.useEffect(() => {
    if (props.isFocused) {
      fetchUser()
    }
  }, [props.isFocused])

  const fetchUser = async () => {
    try {
      const u = await SecureStore.getItemAsync("aumo")
      const res = await axios.get(BACKEND_URL + "/me", {
        headers: {
          Cookie: u
        },
        withCredentials: true
      })
      setUser(res.data)
    } catch (e) {}
    setLoading(false)
  }

  const logout = async () => {
    try {
      await SecureStore.deleteItemAsync("aumo")
    } catch (e) {}
    props.navigation.navigate("LogIn")
  }

  return (
    <View style={styles.container}>
      <ScrollView
        style={styles.container}
        contentContainerStyle={styles.contentContainer}
      >
        {loading ? (
          <ActivityIndicator />
        ) : (
          <>
            <View style={styles.userContainer}>
              <Avatar
                source={{ uri: user.avatar }}
                size="large"
                shape="rounded"
              />
              <View style={{ justifyContent: "center", alignItems: "center" }}>
                <Text
                  style={{ fontSize: 20, fontWeight: "900", marginTop: 10 }}
                >
                  {user.name}
                </Text>
                <View style={{ justifyContent: "space-between" }}>
                  <Text
                    style={{
                      fontSize: 15,
                      marginTop: 10
                    }}
                  >
                    Points:
                    {user.points}
                  </Text>
                </View>
              </View>
              <Button
                style={{ marginTop: 10, borderRadius: 10, textAlign: "right" }}
                size="medium"
                status="basic"
                onPress={logout}
              >
                LOGOUT
              </Button>
            </View>
            <Button
              appearance="ghost"
              onPress={() => {
                if (mode === "rewards") {
                  setMode("receipts")
                  return
                }

                setMode("rewards")
              }}
            >
              Show {mode === "receipts" ? "rewards" : "receipts"}
            </Button>

            {mode === "rewards" ? (
              user.orders &&
              user.orders.length > 0 && (
                <ShopItemList
                  numColumns={1}
                  data={user.orders.map(o => ({
                    ...o,
                    buyable: false,
                    image: { uri: o.image }
                  }))}
                />
              )
            ) : (
              <View style={{ flex: 1 }}>
                <ScrollView>
                  <Receipt
                    total="9,49"
                    shopName="ПАЦОНИ ЕООД, РУСЕ"
                    receiptContent={receiptContent}
                    inCash={"10,00"}
                  />
                  <Receipt
                    total="5,60"
                    shopName="Billa"
                    receiptContent={`Сол 1кг x 2: 2,80лв.`}
                    inCash={"20,00"}
                  />
                </ScrollView>
              </View>
            )}
          </>
        )}
      </ScrollView>
    </View>
  )
}

UserScreen.navigationOptions = {
  header: null
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    height: "100%",
    backgroundColor: "#F7F9FC"
  },
  contentContainer: {
    justifyContent: "center",
    height: "100%",
    paddingTop: 30
  },
  userContainer: {
    backgroundColor: "#fff",
    borderRadius: 12,
    alignSelf: "center",
    justifyContent: "center",
    alignItems: "center",
    padding: 25,
    shadowColor: "#000",
    shadowOffset: {
      width: 0,
      height: 6
    },
    shadowOpacity: 0.37,
    shadowRadius: 7.49,
    elevation: 12,
    marginVertical: 30
  }
})

export default withNavigationFocus(UserScreen)
