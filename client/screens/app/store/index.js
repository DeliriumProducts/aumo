import { Button, Layout } from "@ui-kitten/components"
import React from "react"
import { Context } from "../../../context/context"
import Routes from "../../../navigation/routes"

export default ({ navigation }) => {
  const ctx = React.useContext(Context)

  return (
    <Layout>
      <Button
        onPress={() => {
          navigation.navigate(Routes.StoreShop, {
            id: "123123123",
            name: "Billa"
          })
        }}
      >
        go to billa's shop
      </Button>
    </Layout>
  )
}
