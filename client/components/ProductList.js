import { Button, Card, Icon, List, Text } from "@ui-kitten/components"
import React from "react"
import { Dimensions, ImageBackground, View } from "react-native"
export default ({
  onItemPress = () => {},
  onCartPress = () => {},
  products
}) => {
  const renderProductItem = ({ item }) => (
    <Card
      style={{
        flex: 1,
        margin: 8,
        borderRadius: 16,
        maxWidth: Dimensions.get("window").width / 2 - 24
      }}
      appearance="filled"
      header={() => (
        <ImageBackground
          style={{ height: 140 }}
          source={{ uri: item.image }}
          resizeMode="contain"
        />
      )}
      footer={() => (
        <View
          style={{
            flexDirection: "row",
            justifyContent: "space-between",
            alignItems: "center"
          }}
        >
          <View
            style={{
              flexDirection: "row",
              justifyContent: "center",
              alignItems: "center"
            }}
          >
            <Icon name="award-outline" width={17} height={17} fill="#222" />
            <Text
              category="s1"
              style={{
                marginLeft: 3
              }}
            >
              {item.price}
            </Text>
          </View>
          <Button
            style={{ paddingHorizontal: 0 }}
            size="small"
            icon={style => <Icon {...style} name="shopping-cart" />}
            onPress={() => onCartPress(item)}
          />
        </View>
      )}
      onPress={() => onItemPress(item)}
    >
      <Text category="s1">{item.name}</Text>
      {/* <Text appearance="hint" category="c1">
        {info.item.category}
      </Text> */}
    </Card>
  )

  return (
    <List
      contentContainerStyle={{
        paddingHorizontal: 8,
        paddingVertical: 16
      }}
      data={(products.length && products) || products}
      numColumns={2}
      renderItem={renderProductItem}
    />
  )
}
