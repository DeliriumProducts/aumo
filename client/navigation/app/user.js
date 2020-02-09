import { createStackNavigator } from "@react-navigation/stack"
import React from "react"
import UserScreen from "../../screens/app/user"
import Routes from "../routes"

const Stack = createStackNavigator()

export default () => (
  <Stack.Navigator headerMode="none">
    <Stack.Screen
      name={Routes.User}
      component={UserScreen}
      options={{
        tabBarLabel: ({ focused }) => (
          <Text
            style={{ color: focused ? "#083AA4" : "#CCC", textAlign: "center" }}
          >
            My Profile
          </Text>
        )
      }}
    />
  </Stack.Navigator>
)
