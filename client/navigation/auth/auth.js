import {
  createStackNavigator,
  TransitionPresets
} from "@react-navigation/stack"
import React from "react"
import LoginScreen from "../../screens/auth/login"
import RegisterScreen from "../../screens/auth/register"
import Routes from "../routes"

const Stack = createStackNavigator()

export default () => (
  <Stack.Navigator
    headerMode="screen"
    screenOptions={{
      cardStyle: {
        backgroundColor: theme["color-background-main"]
      },
      ...TransitionPresets.SlideFromRightIOS
    }}
  >
    <Stack.Screen name={Routes.Login} component={LoginScreen} />
    <Stack.Screen name={Routes.Register} component={RegisterScreen} />
  </Stack.Navigator>
)
