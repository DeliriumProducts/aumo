import {
  createStackNavigator,
  TransitionPresets
} from "@react-navigation/stack"
import React from "react"
import StoreScreen from "../../screens/app/store/index.js"
import StoreShopProductScreen from "../../screens/app/store/product"
import StoreShopScreen from "../../screens/app/store/shop.js"
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
    <Stack.Screen name={Routes.Store} component={StoreScreen} />
    <Stack.Screen
      name={Routes.StoreShop}
      component={StoreShopScreen}
      options={({ route }) => ({ title: `${route.params.name}'s store` })}
    />
    <Stack.Screen
      name={Routes.StoreShopProduct}
      component={StoreShopProductScreen}
      options={({ route }) => ({ title: `${route.params.name}` })}
    />
  </Stack.Navigator>
)
