import { List } from "@ui-kitten/components"
import React from "react"
import { View } from "react-native"
import Order from "./Order"

export default ({ orders }) => {
  const renderItem = ({ item: order, index }) => (
    <View style={{ margin: 15 }}>
      <Order product={order.product} key={`${order.order_id}${index}`} />
    </View>
  )

  return <List data={orders} renderItem={renderItem} />
}
