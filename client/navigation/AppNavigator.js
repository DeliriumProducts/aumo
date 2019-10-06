import { createAppContainer, createSwitchNavigator } from "react-navigation"

import AuthLoading from "./AuthLoading"
import MainTabNavigator from "./MainTabNavigator"
import AuthNavigator from "./AuthNavigator"

export default createAppContainer(
  createSwitchNavigator({
    AuthLoading: AuthLoading,
    Main: MainTabNavigator,
    Auth: AuthNavigator
  })
)
