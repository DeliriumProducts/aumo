import { Card, Icon, Text } from "@ui-kitten/components"
import React from "react"
import { Image, View } from "react-native"
import theme from "../theme"

export default ({ product }) => (
  <Card
    appearance="filled"
    style={{ borderRadius: 32 }}
    header={() => (
      <>
        <Image
          source={{ uri: product.image }}
          style={{ height: 400 }}
          resizeMode="contain"
        />
        <View
          style={{
            width: "100%",
            flexDirection: "row",
            flexWrap: "wrap",
            justifyContent: "space-between",
            alignItems: "center",
            padding: 20
          }}
        >
          <Text category="h6">{product.name}</Text>
          <View
            style={{
              flexDirection: "row",
              justifyContent: "center",
              alignItems: "center"
            }}
          >
            <Icon
              name="award-outline"
              width={17}
              height={17}
              fill={theme["color-basic-800"]}
            />
            <Text
              category="h6"
              style={{
                fontWeight: "bold",
                marginLeft: 3
              }}
            >
              {product.price}
            </Text>
          </View>
        </View>
      </>
    )}
  >
    <Text appearance="hint">{product.description}</Text>
  </Card>
)
