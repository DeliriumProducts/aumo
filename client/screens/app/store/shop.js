import { Layout } from "@ui-kitten/components"
import React from "react"
import ProductList from "../../../components/ProductList"

const products = [
  {
    id: 1,
    name: "Akai 32-inch HD LED LCD",
    price: 500,
    image:
      "https://azcd.harveynorman.com.au/media/catalog/product/cache/21/image/992x558/9df78eab33525d08d6e5fb8d27136e95/a/k/ak3219nf.jpg",
    description:
      "Enjoy watching your favourite movies and shows in stunning HD quality with the Akai 32-inch HD LED LCD Smart TV.",
    stock: 5,
    shop_id: 2
  },
  {
    id: 2,
    name: "CYBERPOWERPC Gamer Master Gaming PC",
    price: 500,
    image:
      "https://images-na.ssl-images-amazon.com/images/I/812kz16Md0L._SX466_.jpg",
    description:
      "Cyber PowerPC Gamer Master series is a line of gaming PCs powered by AMD's newest Ryzen CPU and accompanying AM4 architecture.",
    stock: 5,
    shop_id: 2
  },
  {
    id: 2,
    name: "CYBERPOWERPC Gamer Master Gaming PC",
    price: 500,
    image:
      "https://images-na.ssl-images-amazon.com/images/I/812kz16Md0L._SX466_.jpg",
    description:
      "Cyber PowerPC Gamer Master series is a line of gaming PCs powered by AMD's newest Ryzen CPU and accompanying AM4 architecture.",
    stock: 5,
    shop_id: 2
  },
  {
    id: 2,
    name: "CYBERPOWERPC Gamer Master Gaming PC",
    price: 500,
    image:
      "https://images-na.ssl-images-amazon.com/images/I/812kz16Md0L._SX466_.jpg",
    description:
      "Cyber PowerPC Gamer Master series is a line of gaming PCs powered by AMD's newest Ryzen CPU and accompanying AM4 architecture.",
    stock: 5,
    shop_id: 2
  }
]

export default ({ route }) => {
  const { name, id } = route.params
  return (
    <Layout style={{ height: "100%" }} level="1">
      <ProductList products={products} />
    </Layout>
  )
}
