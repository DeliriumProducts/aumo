import { createStackNavigator } from "react-navigation"
import LoginScreen from "../screens/LoginScreen"

export default AuthStack = createStackNavigator(
  {
    LogIn: LoginScreen
  },
  {
    cardStyle: { backgroundColor: "#F5F7FB" }
  }
)
