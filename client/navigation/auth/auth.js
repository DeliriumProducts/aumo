import { createStackNavigator } from "@react-navigation/stack"
import React from "react"
import LoginScreen from "../../screens/auth/login"
import RegisterScreen from "../../screens/auth/register"
import Routes from "../routes"

const Stack = createStackNavigator()

export default () => (
  <Stack.Navigator headerMode="none">
    <Stack.Screen name={Routes.Login} component={LoginScreen} />
    <Stack.Screen name={Routes.Register} component={RegisterScreen} />
  </Stack.Navigator>
)
