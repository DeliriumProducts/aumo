import { Layout, List } from "@ui-kitten/components"
import React from "react"
import Shop from "../../../components/Shop"
import { Context } from "../../../context/context"
import Routes from "../../../navigation/routes"

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
  },
  {
    id: "alsdjf;alksdj",
    name: "Penny Market",
    image: "https://www.capital.bg/shimg/zx620y348_2611418.jpg"
  },
  {
    id: "a;lskdjflasdjfklasd",
    name: "Billa",
    image:
      "https://www.alpbachtal.at/deskline/infrastruktur/objekte/billa-supermarkt_141846/billa-ag_141849.png"
  }
]

export default ({ navigation }) => {
  const ctx = React.useContext(Context)
  const [loading, setLoading] = React.useState()

  return (
    <Layout style={{ height: "100%", padding: 10 }}>
      <List
        data={shops}
        renderItem={({ item: i }) => (
          <Shop
            key={i.id}
            {...i}
            style={{
              width: "100%",
              height: 240,
              marginVertical: 10
            }}
            onPress={() => {
              navigation.navigate(Routes.StoreShop, {
                id: i.id,
                name: i.name
              })
            }}
          />
        )}
      />
    </Layout>
  )
}
