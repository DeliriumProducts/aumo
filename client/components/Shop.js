import { Text } from "@ui-kitten/components"
import React from "react"
import { ImageBackground } from "react-native"

export default ({ name, onPress, image, style }) => (
  <ImageBackground
    source={{ uri: image }}
    onPress={onPress}
    blurRadius={10}
    style={[style, { justifyContent: "center", alignItems: "center" }]}
    imageStyle={{ borderRadius: 32 }}
  >
    <Text category="h1" style={{ color: "white", fontWeight: "bold" }}>
      {name}
    </Text>
  </ImageBackground>
)

// const Image = styled(RImage)``
