import Head from "next/head"
import styled from "styled-components"
import withAuth from "../hocs/withAuth.js"

const Users = () => (
  <>
    <Head>
      <title>Aumo Users</title>
    </Head>
    <Container>users</Container>
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

export default withAuth(Users)
