import ShopScreen from "../../screens/app/shop"
import { createStackNavigator } from "react-navigation-stack"
import Routes from "../routes"

export default createStackNavigator({
  [Routes.Shop]: ShopScreen
})
