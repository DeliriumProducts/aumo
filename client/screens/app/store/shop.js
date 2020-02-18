import { Layout, Text } from "@ui-kitten/components"
import React from "react"

export default ({ route }) => {
  const { name, id } = route.params
  return (
    <Layout>
      <Text>
        hello from the other side! i'm {name}'s shop with id {id}
      </Text>
    </Layout>
  )
}
