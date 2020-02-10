import { createStackNavigator } from "@react-navigation/stack"
import React from "react"
import ShopScreen from "../../screens/app/shop"
import theme from "../../theme"
import Routes from "../routes"

const Stack = createStackNavigator()

export default () => (
  <Stack.Navigator
    headerMode="none"
    screenOptions={{
      cardStyle: {
        backgroundColor: theme["color-background-main"]
      }
    }}
  >
    <Stack.Screen name={Routes.Shop} component={ShopScreen} />
  </Stack.Navigator>
)
