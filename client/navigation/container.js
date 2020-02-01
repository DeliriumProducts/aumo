import { createAppContainer, createSwitchNavigator } from "react-navigation"
import AuthLoading from "./auth/loading"
import Auth from "./auth/auth"
import App from "./app/app"
import Routes from "./routes"

export default createAppContainer(
  createSwitchNavigator({
    [Routes.Auth]: Auth,
    [Routes.AuthLoading]: AuthLoading,
    [Routes.App]: App
  })
)
