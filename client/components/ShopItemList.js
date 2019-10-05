import React from "react"
import { withStyles } from "react-native-ui-kitten"
import { List } from "react-native-ui-kitten"
import ProductListItem from "./ShopItem"

class ProductListComponent extends React.Component {
  constructor() {
    super(...arguments)
    this.onProductAddPress = index => {
      this.props.onItemAddPress(index)
    }
    this.onProductPress = index => {
      this.props.onItemPress(index)
    }
    this.renderListItemElement = item => {
      const { themedStyle } = this.props
      return (
        <ProductListItem
          style={themedStyle.item}
          activeOpacity={0.75}
          image={item.image}
          name={item.name}
          price={`${item.price} points`}
          onAddPress={this.onProductAddPress}
          onPress={this.onProductPress}
        />
      )
    }
    this.renderItem = info => {
      const { item, index } = info
      const listItemElement = this.renderListItemElement(item)
      return React.cloneElement(listItemElement, { index })
    }
  }
  render() {
    const {
      contentContainerStyle,
      themedStyle,
      data,
      ...restProps
    } = this.props
    return (
      <List
        {...restProps}
        contentContainerStyle={[contentContainerStyle, themedStyle.container]}
        data={data}
        renderItem={this.renderItem}
        numColumns={2}
      />
    )
  }
}
export default ProductList = withStyles(ProductListComponent, theme => ({
  container: {
    width: "100%"
  },
  item: {
    flex: 1,
    marginHorizontal: 8,
    marginVertical: 8,
    backgroundColor: theme["background-basic-color-1"]
  }
}))
