import { Button, Icon, Layout, Text } from "@ui-kitten/components"
import React from "react"
import { ImageBackground, ScrollView, View } from "react-native"
import styled from "styled-components/native"
import theme from "../../../theme"

export default ({ route: { params: receipt } }) => {
  return (
    <Header>
      <ShopImage source={{ uri: receipt.shop.image }} />
      <DetailsContainer level="1">
        <View
          style={{
            flexDirection: "row",
            justifyContent: "space-between",
            flexWrap: "wrap"
          }}
        >
          <View>
            <Text category="h4">{receipt.shop.name}</Text>
          </View>
          <View style={{ alignItems: "center" }}>
            <Price>
              <Icon
                name="award-outline"
                width={25}
                height={25}
                fill={theme["color-basic-800"]}
              />
              <Text category="h4">{receipt.total}</Text>
            </Price>
          </View>
        </View>
        <Content appearance="hint">{receipt.content}</Content>
      </DetailsContainer>
    </Header>
  )
}

const Header = styled(ScrollView)`
  min-height: 100%;
`
const ShopImage = styled(ImageBackground)`
  height: 340px;
  width: 100%;
`
const DetailsContainer = styled(Layout)`
  padding-vertical: 24px;
  padding-horizontal: 16px;
`

const Price = styled(View)`
  flex-direction: row;
  align-items: center;
`

const Content = styled(Text)`
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

const ModalContainer = styled(Layout)`
  justify-content: center;
  align-items: center;
  border-radius: 10px;
  padding: 16px;
`
