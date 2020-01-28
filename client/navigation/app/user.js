import { createStackNavigator } from "react-navigation-stack"
import React from "react"
import { Text } from "react-native"
import Routes from "../routes"
import UserScreen from "../../screens/app/user"

const UserStack = createStackNavigator({
  [Routes.User]: UserScreen
})

UserStack.navigationOptions = {
  tabBarLabel: ({ focused }) => (
    <Text style={{ color: focused ? "#083AA4" : "#CCC", textAlign: "center" }}>
      Me
    </Text>
  )
}

UserStack.path = ""

export default UserStack
