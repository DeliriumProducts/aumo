import React from "react"
import { View } from "react-native"

export default Cirlce = ({ color, size, children }) => {
  return (
    <View
      style={{
        borderRadius: Math.round(size) / 2,
        width: size,
        height: size,
        backgroundColor: color,
        justifyContent: "center",
        alignItems: "center"
      }}
    >
      {children}
    </View>
  )
}
