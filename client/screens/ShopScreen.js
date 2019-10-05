import React from "react"
import ShopItem from "../components/ShopItem"
import { ScrollView, StyleSheet, Text, View } from "react-native"

export default function ShopScreen() {
  return (
    <View style={styles.container}>
      <ScrollView
        style={styles.container}
        contentContainerStyle={styles.contentContainer}
      >
        <Text>Epiiic</Text>
        <ShopItem
          image={{
            uri:
              "https://www.pngtube.com/myfile/detail/153-1532616_amazon-discount-gift-card-amazon-gift-card-png.png"
          }}
          name="Amazon Gift Card"
          price="1000"
        />
      </ScrollView>
    </View>
  )
}

ShopScreen.navigationOptions = {
  header: null
}

const styles = StyleSheet.create({
  container: {
    backgroundColor: "#F7F9FC",
    flex: 1
  },
  contentContainer: {
    justifyContent: "center",
    height: "100%",
    alignItems: "center",
    paddingTop: 30
  }
})
