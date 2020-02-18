import { Text } from "@ui-kitten/components"
import React from "react"

export default ({ route }) => (
  <>
    <Text>{route.params.name}</Text>
    <Text>{route.params.description}</Text>
  </>
)
