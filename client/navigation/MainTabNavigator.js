import React from "react"
import { Platform, Text } from "react-native"
import {
  createStackNavigator,
  createBottomTabNavigator
} from "react-navigation"

import TabBarIcon from "../components/TabBarIcon"
import HomeScreen from "../screens/HomeScreen"
import UserScreen from "../screens/UserScreen"
import ShopScreen from "../screens/ShopScreen"

const config = Platform.select({
  web: { headerMode: "screen" },
  default: {},

  cardStyle: { backgroundColor: "#AAA" }
})

const HomeStack = createStackNavigator(
  {
    Home: HomeScreen
  },
  config
)

HomeStack.navigationOptions = {
  tabBarIcon: ({ focused }) => <TabBarIcon focused={focused} name="home" />,
  tabBarLabel: ({ focused }) => (
    <Text style={{ color: focused ? "#083AA4" : "#CCC", textAlign: "center" }}>
      Home
    </Text>
  )
}

HomeStack.path = ""

const ShopStack = createStackNavigator(
  {
    Shop: ShopScreen
  },
  config
)

ShopStack.navigationOptions = {
  tabBarLabel: ({ focused }) => (
    <Text style={{ color: focused ? "#083AA4" : "#CCC", textAlign: "center" }}>
      Shop
    </Text>
  ),
  tabBarIcon: ({ focused }) => (
    <TabBarIcon focused={focused} name="shoppingcart" />
  )
}

ShopStack.path = ""

const UserStack = createStackNavigator(
  {
    User: UserScreen
  },
  config
)

UserStack.navigationOptions = {
  tabBarLabel: ({ focused }) => (
    <Text style={{ color: focused ? "#083AA4" : "#CCC", textAlign: "center" }}>
      User
    </Text>
  ),
  tabBarIcon: ({ focused }) => <TabBarIcon focused={focused} name="user" />
}

UserStack.path = ""

const tabNavigator = createBottomTabNavigator({
  HomeStack,
  ShopStack,
  UserStack
})

tabNavigator.path = ""

export default tabNavigator
