import React from "react"
import { View } from "react-native"

export default Cirlce = ({ color, size, children, ...rest }) => {
  return (
    <View
      {...rest}
      style={{
        borderRadius: Math.round(size) / 2,
        width: size,
        height: size,
        backgroundColor: "#fff",
        justifyContent: "center",
        shadowColor: "#000",
        shadowOffset: {
          width: 0,
          height: 6
        },
        shadowOpacity: 0.37,
        shadowRadius: 7.49,

        elevation: 12,
        alignItems: "center"
      }}
    >
      {children}
    </View>
  )
}
