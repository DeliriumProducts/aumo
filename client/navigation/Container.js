import { createAppContainer, createSwitchNavigator } from "react-navigation"
import AuthLoading from "./auth/loading"
import Auth from "./auth/auth"
import App from "./app/app"
import Routes from "./routes"

export default createAppContainer(
  createSwitchNavigator({
    [Routes.AuthLoading]: AuthLoading,
    [Routes.Auth]: Auth,
    [Routes.App]: App
  })
)
