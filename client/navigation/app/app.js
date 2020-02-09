import { createBottomTabNavigator } from "@react-navigation/bottom-tabs"
import { Icon } from "@ui-kitten/components"
import React from "react"
import { TabBar } from "react-native-animated-nav-tab-bar"
import Routes from "../routes"
import HomeStack from "./home"
import ShopStack from "./shop"
import UserStack from "./user"

const Tab = createBottomTabNavigator()

export default () => (
  <Tab.Navigator
    initialRouteName={Routes.Home}
    tabBarOptions={{
      activeTintColor: "#083aa4",
      inactiveTintColor: "#222222"
    }}
    tabBar={props => (
      <TabBar
        activeColors={"#083aa4"}
        activeTabBackgrounds={"#cae0fa"}
        {...props}
      />
    )}
  >
    <Tab.Screen
      name={Routes.Home}
      component={HomeStack}
      options={{
        tabBarIcon: ({ focused, tintColor }) => (
          <Icon
            name="home-outline"
            fill={focused ? tintColor : "#222222"}
            width={24}
            height={24}
          />
        )
      }}
    />
    <Tab.Screen
      name={Routes.Shop}
      component={ShopStack}
      options={{
        tabBarIcon: ({ focused, tintColor }) => (
          <Icon
            name="shopping-cart-outline"
            fill={focused ? tintColor : "#222222"}
            width={24}
            height={24}
          />
        )
      }}
    />
    <Tab.Screen
      name={Routes.User}
      component={UserStack}
      options={{
        tabBarIcon: ({ focused, tintColor }) => (
          <Icon
            name="person-outline"
            fill={focused ? tintColor : "#222222"}
            width={24}
            height={24}
          />
        )
      }}
    />
  </Tab.Navigator>
)
