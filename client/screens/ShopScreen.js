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
        <ShopItem
          image={{
            uri:
              "https://www.supermarketnews.com/sites/supermarketnews.com/files/styles/article_featured_retina/public/Lidl_US_store_exterior.png?itok=ygwb_d3h"
          }}
          name="Lidl 5% off your next purchase"
          price="10000 points"
        />
        <ShopItem
          image={{
            uri:
              "https://www.supermarketnews.com/sites/supermarketnews.com/files/styles/article_featured_retina/public/Lidl_US_store_exterior.png?itok=ygwb_d3h"
          }}
          name="Lidl 5% off your next purchase"
          price="10000 points"
        />
        <ShopItem
          image={{
            uri:
              "https://www.supermarketnews.com/sites/supermarketnews.com/files/styles/article_featured_retina/public/Lidl_US_store_exterior.png?itok=ygwb_d3h"
          }}
          name="Lidl 5% off your next purchase"
          price="10000 points"
        />
        <ShopItem
          image={{
            uri:
              "https://www.supermarketnews.com/sites/supermarketnews.com/files/styles/article_featured_retina/public/Lidl_US_store_exterior.png?itok=ygwb_d3h"
          }}
          name="Lidl 5% off your next purchase"
          price="10000 points"
        />
        <ShopItem
          image={{
            uri:
              "https://www.supermarketnews.com/sites/supermarketnews.com/files/styles/article_featured_retina/public/Lidl_US_store_exterior.png?itok=ygwb_d3h"
          }}
          name="Lidl 5% off your next purchase"
          price="10000 points"
        />
        <ShopItem
          image={{
            uri:
              "https://www.supermarketnews.com/sites/supermarketnews.com/files/styles/article_featured_retina/public/Lidl_US_store_exterior.png?itok=ygwb_d3h"
          }}
          name="Lidl 5% off your next purchase"
          price="10000 points"
        />
        <ShopItem
          image={{
            uri:
              "https://www.supermarketnews.com/sites/supermarketnews.com/files/styles/article_featured_retina/public/Lidl_US_store_exterior.png?itok=ygwb_d3h"
          }}
          name="Lidl 5% off your next purchase"
          price="10000 points"
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
    paddingTop: 30,
    paddingHorizontal: 60
  }
})
