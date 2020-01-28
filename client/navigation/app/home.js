import { createStackNavigator } from "react-navigation-stack"
import HomeScreen from "../../screens/app/home"
import Routes from "../routes"

export default createStackNavigator({
  [Routes.Home]: HomeScreen
})
