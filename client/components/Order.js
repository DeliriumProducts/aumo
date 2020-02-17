import { Card, Text } from "@ui-kitten/components"
import React from "react"
import { Image } from "react-native"

export default ({ product }) => (
  <Card
    header={() => (
      <>
        <Image source={{ uri: product.image }} style={{ height: 100 }} />
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
