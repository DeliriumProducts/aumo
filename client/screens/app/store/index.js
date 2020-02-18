import { Layout } from "@ui-kitten/components"
import React from "react"
import Shop from "../../../components/Shop"
import { Context } from "../../../context/context"

const shops = [
  {
    id: "asd;flkajs;df",
    name: "Lidl",
    image:
      "https://www.retaildetail.eu/sites/default/files/news/shutterstock_1367384339.jpg"
  },
  {
    id: "asldf;ljas;dlfkjasdf",
    name: "Paconi",
    image: "https://paconi.net/wp-content/uploads/2014/07/20141215_093143.jpg"
  }
]

export default ({ navigation }) => {
  const ctx = React.useContext(Context)

  return (
    <Layout style={{ height: "100%", padding: 10 }}>
      {shops.map(i => (
        <Shop
          key={i.id}
          {...i}
          style={{
            width: "100%",
            height: 200,
            marginVertical: 10
          }}
        />
      ))}
      {/* <Button
        onPress={() => {
          navigation.navigate(Routes.StoreShop, {
            id: "123123123",
            name: "Billa"
          })
        }}
      >
        go to billa's shop
      </Button> */}
    </Layout>
  )
}
