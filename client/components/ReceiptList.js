import { Avatar, Icon, List, ListItem } from "@ui-kitten/components"
import React from "react"

export default ({ receipts, onItemPress = () => {} }) => {
  const renderItem = ({ item: receipt }) =>
    console.log(receipt) || (
      <ListItem
        style={{ borderRadius: 16, margin: 5 }}
        title={receipt.shop.name}
        description={receipt.total}
        onPress={() => onItemPress(receipt)}
        icon={styles => <Icon name="shopping-bag-outline" {...styles} />}
        accessory={styles => (
          <Avatar
            source={{ uri: receipt.shop.image }}
            {...styles}
            size="small"
          />
        )}
      />
    )

  return <List data={receipts} renderItem={renderItem} />
}
