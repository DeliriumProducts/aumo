import Head from "next/head"
import withAuth from "../hocs/withAuth"
import styled from "styled-components"

export const Products = () => (
  <>
    <Head>
      <title>Aumo</title>
      <link rel="icon" href="/favicon.ico" />
    </Head>
    <Container>
      <Card>sdsd</Card>
      <Card>sdsd</Card>
      <Card>sdsd</Card>
      <Card>sdsd</Card>
      <Card>sdsd</Card>
      <Card>sdsd</Card>
      <Card>sdsd</Card>
      <Card>sdsd</Card>
      <Card>sdsd</Card>
      <Card>sdsd</Card>
      <Card>sdsd</Card>
      <Card>sdsd</Card>
      <Card>sdsd</Card>
      <Card>sdsd</Card>
      <Card>sdsd</Card>
      <Card>sdsd</Card>
    </Container>
  </>
)

const Card = styled.div`
  background-color: #fff;
  width: 200px;
  height: 150px;
  text-align: center;
  padding: 20px;
  border-radius: 30px;
  box-shadow: 0px 2px 2px rgba(0, 0, 0, 0.12);

  & p {
    color: black;
    font-weight: 600;
  }
`
const Container = styled.div`
  margin-top: 190px;
  width: 100%;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  & > div {
    margin: 2rem;
  }
  @media only screen and (max-width: 600px) {
    align-items: center;
    flex-direction: column;
  }
`

export default withAuth(Products)
