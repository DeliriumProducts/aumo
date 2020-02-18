import { TabBar } from "@deliriumproducts/react-native-animated-nav-tab-bar"
import { createBottomTabNavigator } from "@react-navigation/bottom-tabs"
import { TransitionPresets } from "@react-navigation/stack"
import { Icon } from "@ui-kitten/components"
import React from "react"
import Avatar from "../../components/Avatar"
import { Context } from "../../context/context"
import theme from "../../theme"
import Routes from "../routes"
import HomeStack from "./home"
import StoreStack from "./store"
import UserStack from "./user"

const Tab = createBottomTabNavigator()

export default () => {
  const ctx = React.useContext(Context)
  return (
    <Tab.Navigator
      screenOptions={{
        ...TransitionPresets.SlideFromRightIOS
      }}
      initialRouteName={Routes.Home}
      tabBarOptions={{
        activeTintColor: theme["color-primary-500"],
        inactiveTintColor: "#222222"
      }}
      tabBar={props => (
        <TabBar
          activeColors={theme["color-primary-500"]}
          activeTabBackgrounds={theme["color-primary-100"]}
          {...props}
        />
      )}
    >
      <Tab.Screen
        name={Routes.Home}
        component={HomeStack}
        options={{
          tabBarIcon: ({ focused, color }) => (
            <Icon
              name="home-outline"
              fill={focused ? color : "#222222"}
              width={24}
              height={24}
            />
          )
        }}
      />
      <Tab.Screen
        name={Routes.Store}
        component={StoreStack}
        options={{
          tabBarIcon: ({ focused, color }) => (
            <Icon
              name="shopping-cart-outline"
              fill={focused ? color : "#222222"}
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
          tabBarLabel: "Profile",
          tabBarIcon: _ => (
            <Avatar
              source={{ uri: ctx.state.user?.avatar }}
              fallbackSource={require("../../assets/Avatar.png")}
            />
          )
        }}
      />
    </Tab.Navigator>
  )
}
