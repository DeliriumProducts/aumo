import ShopScreen from "../../screens/app/shop"
import React from "react"
import { Text } from "react-native"
import { createStackNavigator } from "react-navigation-stack"
import Routes from "../routes"

const ShopStack = createStackNavigator({
  [Routes.Shop]: ShopScreen
})

ShopStack.navigationOptions = {
  tabBarLabel: ({ focused }) => (
    <Text style={{ color: focused ? "#083AA4" : "#CCC", textAlign: "center" }}>
      Shop
    </Text>
  )
}

ShopStack.path = ""

export default ShopStack
