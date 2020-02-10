import { createStackNavigator } from "@react-navigation/stack"
import React from "react"
import UserScreen from "../../screens/app/user"
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
    <Stack.Screen name={Routes.User} component={UserScreen} />
  </Stack.Navigator>
)
