import { createStackNavigator } from "react-navigation-stack"
import Routes from "../routes"
import UserScreen from "../../screens/app/user"

export default createStackNavigator({
  [Routes.User]: UserScreen
})
