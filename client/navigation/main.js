import { createStackNavigator } from "@react-navigation/stack"
import React from "react"
import App from "./app/app"
import Auth from "./auth/auth"
import Routes from "./routes"

const Stack = createStackNavigator()

export default props => (
  <Stack.Navigator {...props} headerMode="none">
    {console.log(props) || props.isAuthenticated ? (
      <Stack.Screen name={Routes.App} component={App} />
    ) : (
      <Stack.Screen name={Routes.Auth} component={Auth} />
    )}
  </Stack.Navigator>
)
