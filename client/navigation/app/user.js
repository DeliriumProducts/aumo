import { createStackNavigator } from "react-navigation-stack"
import UserScreen from "../../screens/app/user"

export default createStackNavigator({
  [Routes.User]: UserScreen
})
