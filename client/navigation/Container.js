import { createAppContainer, createSwitchNavigator } from "react-navigation"
import AuthLoadingNavigator from "./AuthLoadingNavigator"
import AuthNavigator from "./AuthNavigator"
import AppNavigator from "./AppNavigator"
import Routes from "./routes"

export default createAppContainer(
  createSwitchNavigator({
    [Routes.AuthLoading]: AuthLoadingNavigator,
    [Routes.Auth]: AuthNavigator,
    [Routes.App]: AppNavigator
  })
)
