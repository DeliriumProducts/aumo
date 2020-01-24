import Head from "next/head"
import React from "react"
import { UsersAPI } from "aumo-api"
import UserCard from "../components/UserCard"
import styled from "styled-components"
import withAuth from "../hocs/withAuth.js"
import { BACKEND_URL } from "../config"

const Users = () => {
  const [users, setUsers] = React.useState([])
  React.useEffect(() => {
    ;(async () => {
      const data = await new UsersAPI(BACKEND_URL).getAll()
      setUsers(data)
    })()
  })
  return (
    <>
      <Head>
        <title>Aumo Users</title>
      </Head>
      <Container>users</Container>
    </>
  )
}
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
