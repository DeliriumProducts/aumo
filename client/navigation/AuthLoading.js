import { createStackNavigator } from "react-navigation"
import AuthLoadingScreen from "../screens/AuthLoadingScreen"

export default AuthLoadingStack = createStackNavigator(
  {
    AuthLoading: AuthLoadingScreen
  },
  {
    cardStyle: { backgroundColor: "#AAA" }
  }
)
