import React from "react"
import { View, Image, TouchableOpacity } from "react-native"
import { withStyles, Icon } from "react-native-ui-kitten"
import { Button, Text } from "react-native-ui-kitten"

class ProductListItemComponent extends React.Component {
  constructor() {
    super(...arguments)
    this.onAddToBucket = () => {
      if (this.props.buyable) {
        this.props.onAddPress(this.props.index)
      }
    }
  }
  render() {
    const {
      style,
      themedStyle,
      image,
      name,
      price,
      buyable = true,
      ...restProps
    } = this.props
    return (
      <View {...restProps} style={[themedStyle.container, style]}>
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
            {buyable && (
              <Button
                style={themedStyle.buyButton}
                icon={style => <Icon {...style} name="shopping-cart-outline" />}
                onPress={this.onAddToBucket}
              />
            )}
          </View>
        </View>
      </View>
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
      overflow: "hidden",
      shadowColor: "#000",
      shadowOffset: {
        width: 0,
        height: 6
      },
      shadowOpacity: 0.37,
      shadowRadius: 7.49,

      elevation: 12
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
