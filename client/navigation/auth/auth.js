import { createStackNavigator } from "react-navigation-stack"
import Routes from "../routes"
import LoginScreen from "../../screens/auth/login"
import RegisterScreen from "../../screens/auth/register"

export default AuthStack = createStackNavigator({
  [Routes.LogIn]: LoginScreen,
  [Routes.Register]: RegisterScreen
})
