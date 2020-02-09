import { createStackNavigator } from "@react-navigation/stack"
import React from "react"
import App from "./app/app"
import Auth from "./auth/auth"
import Routes from "./routes"

const Stack = createStackNavigator()

export default props => (
  <Stack.Navigator {...props} headerMode="none">
    <Stack.Screen name={Routes.Auth} component={Auth} />
    <Stack.Screen name={Routes.App} component={App} />
  </Stack.Navigator>
)
