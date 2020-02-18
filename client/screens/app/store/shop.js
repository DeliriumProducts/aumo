import { Layout, Modal, Spinner } from "@ui-kitten/components"
import aumo from "aumo"
import React from "react"
import { Alert } from "react-native"
import styled from "styled-components/native"
import ProductList from "../../../components/ProductList"
import { Context } from "../../../context/context"

export default ({ route }) => {
  const ctx = React.useContext(Context)
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

  const placeOrder = async product => {
    try {
      setLoading(true)
      await aumo.order.placeOrder({
        product_id: product.id
      })
      Alert.alert("Successfull!", "You successfully purchased " + product.name)
      setProducts(products =>
        products.map(p => {
          if (p.id == product.id) {
            return {
              ...p,
              stock: p.stock - 1
            }
          }
          return p
        })
      )
    } catch (error) {
      Alert.alert("Error!", error.response.data.error)
    } finally {
      setLoading(false)
    }
  }

  return (
    <Layout style={{ height: "100%" }} level="1">
      <ProductList
        products={products.map(p => ({
          ...p,
          disabled: p.stock < 1 || ctx.state.user.points < p.price
        }))}
        onCartPress={product => {
          Alert.alert(
            "Purchase confirmation",
            `Would you want to order ${product.name}?`,
            [
              {
                text: "Yes",
                onPress: () => placeOrder(product)
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
      />
      {loading && (
        <Modal onBackdropPress={() => {}} visible={loading}>
          <ModalContainer level="1">
            <Spinner size="giant" />
          </ModalContainer>
        </Modal>
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
