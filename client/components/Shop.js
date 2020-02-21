import { Text } from "@ui-kitten/components"
import React from "react"
import { ImageBackground, TouchableOpacity } from "react-native"

export default ({ name, onPress = () => {}, image, style }) => (
  <TouchableOpacity onPress={onPress}>
    <ImageBackground
      source={{ uri: image }}
      blurRadius={8}
      style={[style, { justifyContent: "center", alignItems: "center" }]}
      imageStyle={{ borderRadius: 16 }}>
      <Text category="h1" style={{ color: "white", fontWeight: "bold" }}>
        {name}
      </Text>
    </ImageBackground>
  </TouchableOpacity>
)

// const Image = styled(RImage)``
