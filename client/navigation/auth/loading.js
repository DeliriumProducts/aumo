import Routes from "./routes"
import { createStackNavigator } from "react-navigation-stack"
import AuthLoadingScreen from "../../screens/auth/loading"

export default createStackNavigator({
  [Routes.AuthLoading]: AuthLoadingScreen
})
