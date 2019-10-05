import axios from "axios"
import ShopItemList from "../components/ShopItemList"
import * as SecureStore from "expo-secure-store"
import React from "react"
import {
  ScrollView,
  StyleSheet,
  Text,
  View,
  ActivityIndicator
} from "react-native"
import { BACKEND_URL } from "../config"
import { Avatar, Button } from "react-native-ui-kitten"
import { withNavigationFocus } from "react-navigation"

const UserScreen = props => {
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
                appearance="ghost"
                state="outline"
                onPress={logout}
              >
                LOGOUT
              </Button>
            </View>
            {user.orders && user.orders.length > 0 && (
              <ShopItemList
                data={user.orders.map(o => ({ ...o, buyable: false }))}
              />
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
    paddingTop: 30,
    alignItems: "center"
  },
  userContainer: {
    backgroundColor: "#fff",
    borderRadius: 12,
    justifyContent: "center",
    alignItems: "center",
    padding: 50,
    shadowColor: "#000",
    shadowOffset: {
      width: 0,
      height: 6
    },
    shadowOpacity: 0.37,
    shadowRadius: 7.49,

    elevation: 12
  }
})

export default withNavigationFocus(UserScreen)
