import React from "react"
import { Platform, Text } from "react-native"
import {
  createStackNavigator,
  createBottomTabNavigator
} from "react-navigation"

import TabBarIcon from "../components/TabBarIcon"
import HomeScreen from "../screens/HomeScreen"
import LinksScreen from "../screens/LinksScreen"
import SettingsScreen from "../screens/SettingsScreen"

const config = Platform.select({
  web: { headerMode: "screen" },
  default: {}
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

const LinksStack = createStackNavigator(
  {
    Links: LinksScreen
  },
  config
)

LinksStack.navigationOptions = {
  tabBarLabel: "Links",
  tabBarIcon: ({ focused }) => (
    <TabBarIcon focused={focused} name="shoppingcart" />
  )
}

LinksStack.path = ""

const SettingsStack = createStackNavigator(
  {
    Settings: SettingsScreen
  },
  config
)

SettingsStack.navigationOptions = {
  tabBarLabel: "Settings",
  tabBarIcon: ({ focused }) => <TabBarIcon focused={focused} name="user" />
}

SettingsStack.path = ""

const tabNavigator = createBottomTabNavigator({
  HomeStack,
  LinksStack,
  SettingsStack
})

tabNavigator.path = ""

export default tabNavigator
