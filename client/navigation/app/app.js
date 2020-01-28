import { createBottomTabNavigator } from "react-navigation-tabs"

import HomeStack from "./home"
import UserStack from "./user"
import ShopStack from "./shop"

export default createBottomTabNavigator({
  HomeStack,
  ShopStack,
  UserStack
})
