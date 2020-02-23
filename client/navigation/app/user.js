import {
  createStackNavigator,
  TransitionPresets
} from "@react-navigation/stack"
import React from "react"
import UserScreen from "../../screens/app/user"
import EditUserScreen from "../../screens/app/user/edit"
import ReceiptDetailsScreen from "../../screens/app/user/receipt"
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
    <Stack.Screen name={Routes.UserEdit} component={EditUserScreen} />
    <Stack.Screen
      name={Routes.ReceiptDetails}
      component={ReceiptDetailsScreen}
    />
  </Stack.Navigator>
)
