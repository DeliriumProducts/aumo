import Head from "next/head"
import React from "react"
import { UserAPI } from "aumo-api"
import styled from "styled-components"
import withAuth from "../hocs/withAuth.js"
import { BACKEND_URL } from "../config"

const Users = () => {
  const [users, setUsers] = React.useState([])
  React.useEffect(() => {
    ;(async () => {
      const data = await new UserAPI(BACKEND_URL).getAll()
      setUsers(data.data)
    })()
  }, [])
  console.log(users)
  return (
    <>
      <Head>
        <title>Aumo Users</title>
      </Head>
      <Container>
        {users &&
          users.length > 0 &&
          users.map(u => (
            <UserCard
              key={u.id}
              name={u.name}
              email={u.email}
              avatar={u.avatar}
            />
          ))}
      </Container>
    </>
  )
}

const UserCard = ({ name, email, avatar, onDelete }) => {
  return (
    <UserCardContainer>
      <img src={avatar} />
      <NameContainer>
        {name}
        {email}
      </NameContainer>
    </UserCardContainer>
  )
}

const UserCardContainer = styled.div`
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: white;
  margin-bottom: 1rem;
  border-radius: 30px;
  box-shadow: rgba(0, 0, 0, 0.31) 0px 20px 24px -18px;
`

const NameContainer = styled.div`
  display: flex;
  width: 100%;
  flex-direction: column;
  justify-content: center;
  align-items: center;
`

const Container = styled.div`
  height: 100%;
  width: 100%;
  display: flex;
  justify-content: center;
  padding: 10rem;
  align-items: center;
  flex-direction: column;
  @media only screen and (max-width: 600px) {
    flex-direction: column;
  }
`

export default withAuth(Users)
