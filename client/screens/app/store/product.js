import { Button, Icon, Layout, Text } from "@ui-kitten/components"
import React from "react"
import { ImageBackground, View } from "react-native"
import styled from "styled-components/native"

export default ({
  route: {
    params: { product, shop }
  }
}) => {
  const onPress = async () => {}

  return (
    <Header>
      <ProductImage source={{ uri: product.image }} resizeMode="contain" />
      <DetailsContainer level="1">
        <View
          style={{
            flexDirection: "row",
            justifyContent: "space-between",
            flexWrap: "wrap"
          }}
        >
          <View>
            <Text category="h6">{product.name}</Text>
            <Shop appearance="hint" category="p2">
              {shop.name}
            </Shop>
          </View>
          <Price>
            <Icon name="award-outline" width={30} height={30} fill="#8f9bb3" />
            <Text category="h4">{product.price}</Text>
          </Price>
        </View>
        <Description appearance="hint">{product.description}</Description>
        <ActionContainer>
          <ActionButton size="giant" onPress={onPress}>
            BUY
          </ActionButton>
        </ActionContainer>
      </DetailsContainer>
    </Header>
  )
}

const Header = styled(Layout)`
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
  margin-horizontal: 8px;
`
