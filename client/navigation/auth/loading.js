import { createStackNavigator } from "react-navigation-stack"
import Routes from "../routes"
import AuthLoadingScreen from "../../screens/auth/loading"

export default createStackNavigator({
  [Routes.AuthLoading]: AuthLoadingScreen
})
