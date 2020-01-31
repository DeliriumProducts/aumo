import React from "react"
import { ActivityIndicator, StatusBar, View } from "react-native"
import { BACKEND_URL } from "../../config/index.js"
import { AuthAPI } from "aumo-api"
import Routes from "../../navigation/routes"

export default ({ navigation }) => {
  React.useEffect(() => {
    ;(async () => {
      try {
        await new AuthAPI(BACKEND_URL).me()
        navigation.navigate(Routes.App)
      } catch (e) {
        console.warn(e)
        if (e.response.status === 401) {
          navigation.navigate(Routes.Auth)
        }
      }
    })()
  }, [])

  return (
    <View>
      <ActivityIndicator />
      <StatusBar barStyle="default" />
    </View>
  )
}
