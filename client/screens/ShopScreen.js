import React from "react"
import { BACKEND_URL } from "../config"
import { withNavigationFocus } from "react-navigation"
import axios from "axios"
import ShopItemList from "../components/ShopItemList"
import {
  ScrollView,
  StyleSheet,
  Text,
  View,
  ActivityIndicator
} from "react-native"

function ShopScreen(props) {
  const [items, setItems] = React.useState(false)
  const [loading, setLoading] = React.useState(true)

  React.useEffect(() => {
    ;(async () => {
      const res = await axios.get(BACKEND_URL + "/shop")

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
        {loading ? <ActivityIndicator /> : <ShopItemList data={items} />}
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
