import Head from "next/head"
import { withAuth } from "../hocs/withAuth"

export const Products = () => (
  <>
    <Head>
      <title>Aumo</title>
      <link rel="icon" href="/favicon.ico" />
    </Head>
    Products
  </>
)

export default withAuth(Products)
