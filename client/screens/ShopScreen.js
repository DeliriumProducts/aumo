import React from "react"
import { ProductAPI, OrderAPI } from "../api"
import { BACKEND_URL } from "../config"
import { withNavigationFocus } from "react-navigation"
import axios from "axios"
import ShopItemList from "../components/ShopItemList"
import {
  ScrollView,
  StyleSheet,
  Alert,
  Text,
  View,
  ActivityIndicator
} from "react-native"

function ShopScreen(props) {
  const [items, setItems] = React.useState(false)
  const [loading, setLoading] = React.useState(true)

  React.useEffect(() => {
    ;(async () => {
      const res = await ProductAPI.getAll()

      setItems(res.data)
      setLoading(false)
    })()
  }, [props.isFocused])

  return (
    <View style={styles.container}>
      <ScrollView
        style={styles.container}
        contentContainerStyle={styles.contentContainer}
      >
        {loading ? (
          <ActivityIndicator />
        ) : (
          <ShopItemList
            onItemAddPress={async idx => {
              console.log("here")
              const shopItem = items[idx]
              Alert.alert(
                "Purchase confirmation",
                "Would you want to buy " + shopItem.name,
                [
                  {
                    text: "Yes",
                    onPress: async () => {
                      try {
                        const res = await OrderAPI.placeOrder(shopItem.id)
                        if (res.status === 200) {
                          Alert.alert(
                            "Successfull!",
                            "You successfully purchased " + shopItem.name
                          )
                        }
                      } catch (e) {
                        if (e.response.status === 500) {
                          Alert.alert("Error!", "Some error occured!")
                        }
                      }
                    }
                  },
                  {
                    text: "Cancel",
                    onPress: () => {},
                    style: "cancel"
                  }
                ],
                { cancelable: true }
              )
            }}
            data={items.map(i => ({
              ...i,
              image: { uri: i.image },
              buyable: true
            }))}
          />
        )}
      </ScrollView>
    </View>
  )
}

ShopScreen.navigationOptions = {
  header: null
}

const styles = StyleSheet.create({
  container: {
    backgroundColor: "#F7F9FC",
    flex: 1
  },
  contentContainer: {
    justifyContent: "center",
    height: "100%",
    paddingTop: 30
  }
})

export default withNavigationFocus(ShopScreen)
