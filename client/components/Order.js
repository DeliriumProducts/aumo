import { Card, Text } from "@ui-kitten/components"
import React from "react"
import { Image } from "react-native"

export default ({ product }) => (
  <Card
    appearance="filled"
    style={{ borderRadius: 32 }}
    header={() => (
      <>
        <Image source={{ uri: product.image }} style={{ height: 400 }} />
        <Text
          category="h6"
          style={{ marginHorizontal: 24, marginVertical: 16 }}
        >
          {product.name}
        </Text>
      </>
    )}
  >
    <Text>{product.description}</Text>
  </Card>
)
