import React from "react"
import { View, Image, TouchableOpacity } from "react-native"
import { withStyles, Icon } from "react-native-ui-kitten"
import { Button, Text } from "react-native-ui-kitten"

class ProductListItemComponent extends React.Component {
  constructor() {
    super(...arguments)
    this.onPress = () => {
      this.props.onPress(this.props.index)
    }
    this.onAddToBucket = () => {
      this.props.onAddPress(this.props.index)
    }
  }
  render() {
    const { style, themedStyle, image, name, price, ...restProps } = this.props
    return (
      <TouchableOpacity
        {...restProps}
        style={[themedStyle.container, style]}
        onPress={this.onPress}
      >
        <Image style={themedStyle.image} source={image} />
        <View style={themedStyle.infoContainer}>
          <View>
            <Text style={themedStyle.nameLabel} category="s1">
              {name}
            </Text>
          </View>
          <View style={themedStyle.priceContainer}>
            <Text style={themedStyle.costLabel} category="s1">
              {price}
            </Text>
            <Button
              style={themedStyle.buyButton}
              icon={style => <Icon {...style} name="shopping-cart-outline" />}
              onPress={this.onAddToBucket}
            />
          </View>
        </View>
      </TouchableOpacity>
    )
  }
}
export default ProductListItem = withStyles(
  ProductListItemComponent,
  theme => ({
    container: {
      minHeight: 272,
      borderRadius: 12,
      backgroundColor: "#fff",
      overflow: "hidden"
    },
    infoContainer: {
      flex: 1,
      padding: 16,
      justifyContent: "space-between"
    },
    priceContainer: {
      flexDirection: "row",
      alignItems: "center",
      justifyContent: "space-between"
    },
    image: {
      flex: 1,
      width: null,
      height: 140
    },

    buyButton: {
      width: 32,
      height: 32
    }
  })
)
