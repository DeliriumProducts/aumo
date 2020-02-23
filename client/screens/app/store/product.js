import {
  Button,
  Icon,
  Layout,
  Modal,
  Spinner,
  Text
} from "@ui-kitten/components"
import { Context } from "../../../context/context"
import aumo from "aumo"
import React from "react"
import { Alert, ImageBackground, ScrollView, View } from "react-native"
import styled from "styled-components/native"
import theme from "../../../theme"

export default ({
  route: {
    params: {
      product: { id: pID },
      shop
    }
  }
}) => {
  const ctx = React.useContext(Context)
  const [product, setProduct] = React.useState(null)
  const [loading, setLoading] = React.useState(true)

  const fetchProduct = async () => {
    try {
      const data = await aumo.shop.getProduct(shop.id, pID)
      setProduct({
        ...data,
        disabled: data.stock < 1 || ctx.state?.user?.points < data.price
      })
    } catch (error) {
      console.warn(error)
    }
  }

  React.useEffect(() => {
    ;(async () => {
      await fetchProduct()
      setLoading(false)
    })()
  }, [])

  const onPress = async product => {
    try {
      setLoading(true)
      await aumo.order.placeOrder({
        product_id: product.id
      })
      Alert.alert("Successfull!", "You successfully purchased " + product.name)
    } catch (error) {
      Alert.alert("Error!", error.response.data.error)
    } finally {
      setLoading(false)
      fetchProduct()
    }
  }

  return (
    <Header>
      <ProductImage source={{ uri: product?.image }} resizeMode="contain" />
      <DetailsContainer level="1">
        <View
          style={{
            flexDirection: "row",
            justifyContent: "space-between",
            flexWrap: "wrap"
          }}
        >
          <View>
            <Text category="h4">{product?.name}</Text>
            <Shop appearance="hint" category="p2">
              provided by {shop.name}
            </Shop>
          </View>
          <View style={{ alignItems: "center" }}>
            <Price>
              <Icon
                name="award-outline"
                width={25}
                height={25}
                fill={theme["color-basic-800"]}
              />
              <Text category="h4">{product?.price}</Text>
            </Price>
            <Shop
              appearance="hint"
              category="p2"
              status={product?.stock < 1 ? "danger" : ""}
            >
              {product?.stock} in stock
            </Shop>
          </View>
        </View>
        <Description appearance="hint">{product?.description}</Description>
        <ActionContainer>
          <ActionButton
            size="giant"
            disabled={product?.disabled}
            onPress={() => {
              Alert.alert(
                "Purchase confirmation",
                `Would you want to order ${product.name}?`,
                [
                  {
                    text: "Yes",
                    onPress: () => onPress(product)
                  },
                  {
                    text: "Cancel",
                    onPress: () => {},
                    style: "cancel"
                  }
                ],
                { cancelable: true }
              )
            }}
            icon={style => <Icon {...style} name="shopping-cart" />}
          >
            BUY
          </ActionButton>
        </ActionContainer>
      </DetailsContainer>
      {loading && (
        <Modal onBackdropPress={() => {}} visible={loading}>
          <ModalContainer level="1">
            <Spinner size="giant" />
          </ModalContainer>
        </Modal>
      )}
    </Header>
  )
}

const Header = styled(ScrollView)`
  min-height: 100%;
`
const ProductImage = styled(ImageBackground)`
  height: 340px;
  width: 100%;
`
const DetailsContainer = styled(Layout)`
  padding-vertical: 24px;
  padding-horizontal: 16px;
`

const Shop = styled(Text)`
  margin-top: 4px;
`

const Price = styled(View)`
  flex-direction: row;
  align-items: center;
`

const Description = styled(Text)`
  margin-vertical: 16px;
`

const ActionContainer = styled(View)`
  flex-direction: row;
  margin-horizontal: -8px;
  margin-top: 24px;
`

const ActionButton = styled(Button)`
  flex: 1;
  border-radius: 10px;
  margin-horizontal: 8px;
`

const ModalContainer = styled(Layout)`
  justify-content: center;
  align-items: center;
  border-radius: 10px;
  padding: 16px;
`
