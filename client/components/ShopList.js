import { List } from "@ui-kitten/components"
import React from "react"
import Shop from "./Shop"

export default ({ shops, onShopPress }) => {
  const renderItem = ({ item: i }) => (
    <Shop
      key={i.id}
      {...i}
      style={{
        width: "100%",
        height: 240,
        marginVertical: 10
      }}
      onPress={() => onShopPress(i)}
    />
  )

  return (
    <List
      style={{
        padding: 20
      }}
      data={shops}
      renderItem={renderItem}
    />
  )
}
