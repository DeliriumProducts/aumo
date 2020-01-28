import { Button } from "@ui-kitten/components"
import Routes from "../../navigation/routes"
import React from "react"

export default ({ navigation }) => {
  return (
    <Button
      onPress={() => {
        navigation.navigate(Routes.App)
      }}
    >
      Login!
    </Button>
  )
}
