import { createStackNavigator } from "react-navigation-stack"
import HomeScreen from "../../screens/app/home"
import { Text } from "react-native"
import Routes from "../routes"
import React from "react"

const HomeStack = createStackNavigator({
  [Routes.Home]: HomeScreen
})

HomeStack.navigationOptions = {
  tabBarLabel: ({ focused }) => (
    <Text style={{ color: focused ? "#083AA4" : "#CCC", textAlign: "center" }}>
      Home
    </Text>
  )
}

HomeStack.path = ""

export default HomeStack
