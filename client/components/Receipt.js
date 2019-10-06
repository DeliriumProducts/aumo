import React from "react"
import { View, Text, StyleSheet } from "react-native"

export default Receipt = ({ inCash, shopName, receiptContent, total }) => (
  <View style={styles.receiptContainer}>
    <View style={{ justifyContent: "center", alignItems: "center" }}>
      <Text
        style={{
          fontSize: 20,
          fontWeight: "400",
          marginTop: 10
        }}
      >
        {shopName}
      </Text>
      <View style={{ justifyContent: "space-between" }}>
        <Text
          style={{
            fontSize: 15,
            marginTop: 10
          }}
        >
          {receiptContent}
        </Text>
      </View>
      <Text
        style={{
          fontSize: 20,
          fontWeight: "400",
          marginTop: 20
        }}
      >
        ОБЩА СУМА:{" "}
        <Text
          style={{
            fontSize: 20,
            fontWeight: "600",
            marginTop: 10
          }}
        >
          {total}лв.
        </Text>
      </Text>
      <Text
        style={{
          fontSize: 20,
          fontWeight: "400",
          marginTop: 20
        }}
      >
        В БРОЙ{" "}
        <Text
          style={{
            fontSize: 20,
            fontWeight: "600",
            marginTop: 10
          }}
        >
          {inCash}лв.
        </Text>
      </Text>
    </View>
  </View>
)

const styles = StyleSheet.create({
  receiptContainer: {
    backgroundColor: "#fff",
    borderRadius: 12,
    alignSelf: "center",
    justifyContent: "center",
    alignItems: "center",
    padding: 25,
    shadowColor: "#000",
    shadowOffset: {
      width: 0,
      height: 6
    },
    shadowOpacity: 0.37,
    shadowRadius: 7.49,
    elevation: 12,
    marginVertical: 30
  }
})
