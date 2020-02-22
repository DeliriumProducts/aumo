import { Layout, Modal, Spinner } from "@ui-kitten/components"
import aumo from "aumo"
import React from "react"
import { Alert, ActionSheetIOS } from "react-native"
import { actions } from "../../../context/providers/provider"
import styled from "styled-components/native"
import ProductList from "../../../components/ProductList"
import { Context } from "../../../context/context"
import Routes from "../../../navigation/routes"

export default ({ route, navigation }) => {
  const ctx = React.useContext(Context)
  const [products, setProducts] = React.useState([])
  const [loading, setLoading] = React.useState(true)

  const fetchProducts = async () => {
    try {
      const response = await aumo.shop.getAllProductsByShop(route.params.id)
      setProducts(response)
    } catch (error) {
      console.warn(error)
    }
  }

  const fetchUser = async () => {
    try {
      const user = await aumo.auth.me()
      ctx.dispatch({ type: actions.SET_USER, payload: user })
    } catch (error) {}
  }

  React.useEffect(() => {
    ;(async () => {
      await fetchProducts()
      setLoading(false)
    })()
  }, [])

  const placeOrder = async product => {
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
      fetchProducts()
      fetchUser()
    }
  }

  return (
    <Layout style={{ height: "100%" }} level="1">
      <ProductList
        products={products.map(p => ({
          ...p,
          disabled: p.stock < 1 || ctx.state.user.points < p.price
        }))}
        onItemPress={product => {
          navigation.navigate(Routes.ProductDetails, {
            product,
            shop: route.params
          })
        }}
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
