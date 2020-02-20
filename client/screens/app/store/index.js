import { Layout, Spinner } from "@ui-kitten/components"
import aumo from "aumo"
import React from "react"
import styled from "styled-components/native"
import ShopList from "../../../components/ShopList"
import { Context } from "../../../context/context"
import Routes from "../../../navigation/routes"

export default ({ navigation }) => {
  const ctx = React.useContext(Context)
  const [loading, setLoading] = React.useState(true)
  const [shops, setShops] = React.useState([])

  React.useEffect(() => {
    ;(async () => {
      try {
        const response = await aumo.shop.getAllShops()
        setShops(response)
      } catch (error) {
        console.warn(error)
      } finally {
        setLoading(false)
      }
    })()
  }, [])

  return (
    <Layout style={{ height: "100%" }} level="2">
      <ShopList
        shops={shops}
        onShopPress={shop => {
          navigation.navigate(Routes.StoreShop, {
            id: shop.id,
            name: shop.name
          })
        }}
      />
      {loading && (
        <Layout
          level="2"
          style={{
            height: "100%",
            width: "100%",
            position: "absolute",
            justifyContent: "center",
            alignItems: "center"
          }}
        >
          <ModalContainer level="1">
            <Spinner size="giant" />
          </ModalContainer>
        </Layout>
      )}
    </Layout>
  )
}

const ModalContainer = styled(Layout)`
  justify-content: center;
  align-items: center;
  border-radius: 10px;
  padding: 16px;
`
