import {
  createStackNavigator,
  TransitionPresets
} from "@react-navigation/stack"
import React from "react"
import HomeScreen from "../../screens/app/home"
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
    <Stack.Screen name={Routes.Home} component={HomeScreen} />
  </Stack.Navigator>
)
