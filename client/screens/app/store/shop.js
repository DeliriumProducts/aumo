import { Layout, Spinner } from "@ui-kitten/components"
import aumo from "aumo"
import React from "react"
import styled from "styled-components/native"
import ProductList from "../../../components/ProductList"

export default ({ route }) => {
  const [products, setProducts] = React.useState([])
  const [loading, setLoading] = React.useState(true)

  React.useEffect(() => {
    ;(async () => {
      try {
        const response = await aumo.shop.getShop(route.params.id)
        setProducts(response.products)
      } catch (error) {
        console.warn(error)
      } finally {
        setLoading(false)
      }
    })()
  }, [])

  return (
    <Layout style={{ height: "100%" }} level="1">
      <ProductList products={products} />
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
