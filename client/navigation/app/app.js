import { createBottomTabNavigator } from "@react-navigation/bottom-tabs"
import React from "react"
import Routes from "../routes"
import HomeStack from "./home"
import ShopStack from "./shop"
import UserStack from "./user"

const Tab = createBottomTabNavigator()

export default () => (
  <Tab.Navigator initialRouteName={Routes.Home}>
    <Tab.Screen name={Routes.Home} component={HomeStack} />
    <Tab.Screen name={Routes.Shop} component={ShopStack} />
    <Tab.Screen name={Routes.User} component={UserStack} />
  </Tab.Navigator>
)
