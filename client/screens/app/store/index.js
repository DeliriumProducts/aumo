import { Layout, Spinner } from "@ui-kitten/components"
import aumo from "aumo"
import Modal from "../../../components/Modal"
import React from "react"
import styled from "styled-components/native"
import ShopList from "../../../components/ShopList"
import Routes from "../../../navigation/routes"

export default ({ navigation }) => {
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
          navigation.navigate(Routes.StoreShop, shop)
        }}
      />
      <Modal visible={loading}>{loading && <Spinner size="giant" />}</Modal>
    </Layout>
  )
}

const ModalContainer = styled(Layout)`
  justify-content: center;
  align-items: center;
  border-radius: 10px;
  padding: 16px;
`
