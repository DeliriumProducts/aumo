import { createStackNavigator } from "react-navigation"
import LoginScreen from "../screens/LoginScreen"
import RegisterScreen from "../screens/RegisterScreen"

export default AuthStack = createStackNavigator(
  {
    LogIn: LoginScreen,
    Register: RegisterScreen
  },
  {
    cardStyle: { backgroundColor: "#FBFBFB" }
  }
)
