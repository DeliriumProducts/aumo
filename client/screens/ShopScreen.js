import React from "react"
import ShopItemList from "../components/ShopItemList"
import { ScrollView, StyleSheet, Text, View } from "react-native"

export default function ShopScreen() {
  return (
    <View style={styles.container}>
      <ScrollView
        style={styles.container}
        contentContainerStyle={styles.contentContainer}
      >
        <ShopItemList
          data={[
            {
              name: "Lidl 5% off coupon",
              price: 50,
              image: {
                uri:
                  "https://www.supermarketnews.com/sites/supermarketnews.com/files/styles/article_featured_retina/public/Lidl_US_store_exterior.png?itok=ygwb_d3h"
              }
            },
            {
              name: "Billa 15% off coupon",
              price: 150,
              image: {
                uri:
                  "https://www.agrobusiness.bg/images/cache/68097e541d2f6ba1cc10889632311274_w744_h500_cp.jpg"
              }
            },
            {
              name: "Technopolis 3% off coupon",
              price: 30,
              image: {
                uri:
                  "https://www.technopolis.bg/medias/sys_master/h78/h38/11382252896286.jpg"
              }
            },
            {
              name: "MediaMarket 40% off coupon",
              price: 350,
              image: {
                uri:
                  "https://s9783.pcdn.co/wp-content/uploads/2017/04/MediaMarkt-Sweden-Store_small.jpg"
              }
            },
            {
              name: "Lidl 5% off coupon",
              price: 50,
              image: {
                uri:
                  "https://www.supermarketnews.com/sites/supermarketnews.com/files/styles/article_featured_retina/public/Lidl_US_store_exterior.png?itok=ygwb_d3h"
              }
            },
            {
              name: "Billa 15% off coupon",
              price: 150,
              image: {
                uri:
                  "https://www.agrobusiness.bg/images/cache/68097e541d2f6ba1cc10889632311274_w744_h500_cp.jpg"
              }
            },
            {
              name: "Technopolis 3% off coupon",
              price: 30,
              image: {
                uri:
                  "https://www.technopolis.bg/medias/sys_master/h78/h38/11382252896286.jpg"
              }
            },
            {
              name: "MediaMarket 40% off coupon",
              price: 350,
              image: {
                uri:
                  "https://s9783.pcdn.co/wp-content/uploads/2017/04/MediaMarkt-Sweden-Store_small.jpg"
              }
            }
          ]}
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
    paddingTop: 30
  }
})
