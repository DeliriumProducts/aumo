import {
  createStackNavigator,
  TransitionPresets
} from "@react-navigation/stack"
import React from "react"
import UserScreen from "../../screens/app/user"
import theme from "../../theme"
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
    <Stack.Screen name={Routes.User} component={UserScreen} />
  </Stack.Navigator>
)
