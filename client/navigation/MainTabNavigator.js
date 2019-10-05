import React from "react"
import { Platform } from "react-native"
import {
  createStackNavigator,
  createBottomTabNavigator
} from "react-navigation"

import TabBarIcon from "../components/TabBarIcon"
import HomeScreen from "../screens/HomeScreen"
import SettingsScreen from "../screens/SettingsScreen"
import ShopScreen from "../screens/ShopScreen"

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
  tabBarLabel: "Home",
  tabBarIcon: ({ focused }) => <TabBarIcon focused={focused} name="home" />
}

HomeStack.path = ""

const ShopStack = createStackNavigator(
  {
    Shop: ShopScreen
  },
  config
)

ShopStack.navigationOptions = {
  tabBarLabel: "Shop",
  tabBarIcon: ({ focused }) => (
    <TabBarIcon focused={focused} name="shoppingcart" />
  )
}

ShopStack.path = ""

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
  ShopStack,
  SettingsStack
})

tabNavigator.path = ""

export default tabNavigator
