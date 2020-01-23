import Head from "next/head"
import withAuth from "../hocs/withAuth"
import styled from "styled-components"

export const Products = () => (
  <>
    <Head>
      <title>Aumo</title>
      <link rel="icon" href="/favicon.ico" />
    </Head>
    <Container>Products</Container>
  </>
)
const Container = styled.div`
  height: 100%;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  @media only screen and (max-width: 600px) {
    flex-direction: column;
  }
`

export default withAuth(Products)
