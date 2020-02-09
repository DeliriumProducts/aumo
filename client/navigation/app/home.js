import { createStackNavigator } from "@react-navigation/stack"
import React from "react"
import HomeScreen from "../../screens/app/home"
import Routes from "../routes"

const Stack = createStackNavigator()

export default () => (
  <Stack.Navigator headerMode="none">
    <Stack.Screen
      name={Routes.Home}
      component={HomeScreen}
      options={{
        tabBarLabel: ({ focused }) => (
          <Text
            style={{ color: focused ? "#083AA4" : "#CCC", textAlign: "center" }}
          >
            Home
          </Text>
        )
      }}
    />
  </Stack.Navigator>
)
