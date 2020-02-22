import { List } from "@ui-kitten/components"
import React from "react"
import Product from "./Product"

export default ({
  onItemPress = () => {},
  onCartPress = () => {},
  products
}) => {
  const renderItem = ({ item }) => (
    <Product item={item} onItemPress={onItemPress} onCartPress={onCartPress} />
  )

  return (
    <List
      contentContainerStyle={{
        paddingHorizontal: 8,
        paddingVertical: 16
      }}
      data={(products.length && products) || products}
      numColumns={2}
      renderItem={renderItem}
    />
  )
}
