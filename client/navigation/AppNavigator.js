import { createStackNavigator } from "react-navigation-stack"
import { createBottomTabNavigator } from "react-navigation-tabs"
import React from "react"
import HomeScreen from "../screens/app/home"
import ShopScreen from "../screens/app/shop"
import UserScreen from "../screens/app/user"
import Routes from "./routes"

const HomeStack = createStackNavigator({
  [Routes.Home]: HomeScreen
})

const ShopStack = createStackNavigator({
  [Routes.Shop]: ShopScreen
})

const UserStack = createStackNavigator({
  [Routes.User]: UserScreen
})

export default createBottomTabNavigator({
  HomeStack,
  ShopStack,
  UserStack
})
