import { createStackNavigator } from "@react-navigation/stack"
import React from "react"
import ShopScreen from "../../screens/app/shop"
import Routes from "../routes"

const Stack = createStackNavigator()

export default () => (
  <Stack.Navigator headerMode="none">
    <Stack.Screen name={Routes.Shop} component={ShopScreen} />
  </Stack.Navigator>
)
