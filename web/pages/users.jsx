import Head from "next/head"
import { Card, Button, Avatar } from "antd"
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

  const showUser = (e, id) => {
    console.log(id)
  }

  const deleteUser = (e, id) => {
    console.log(id)
  }

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
              id={u.id}
              key={u.id}
              name={u.name}
              onClick={showUser}
              onDelete={deleteUser}
              email={u.email}
              avatar={u.avatar}
            />
          ))}
      </Container>
    </>
  )
}

const UserCard = ({ id, name, email, avatar, onDelete, onClick }) => {
  return (
    <UserCardContainer
      hoverable
      onClick={e => {
        onClick(e, id)
      }}
    >
      <div>
        <Avatar src={avatar} size={80} key={id} className="avatar" />
      </div>
      <NameContainer>
        <h1>{name}</h1>
        <h2>{email}</h2>
      </NameContainer>
      <Filler />
      <Button
        type="danger"
        icon="delete"
        size="large"
        onClick={e => {
          e.stopPropagation()
          onDelete(e, id)
        }}
        style={{
          right: 10
        }}
      >
        Delete
      </Button>
    </UserCardContainer>
  )
}

const Filler = styled.div`
  width: 100%;
`

const UserCardContainer = styled(Card)`
  margin-top: 8px;
  border-radius: 20px;
  width: 100%;
  border: none;
  text-align: center;
  display: flex;
  flex-direction: row;
  padding: 0.5rem;
  height: 8rem;
  box-shadow: rgba(0, 0, 0, 0.31) 0px 20px 24px -18px;

  .ant-card-body {
    width: 100%;
    display: flex;
    padding-top: 8px;
    padding-bottom: 8px;
    align-items: center;
    padding-left: 0;
    padding-right: 0;
    justify-content: center;
  }

  .ant-btn-group {
    min-width: 8rem;
  }

  .ant-avatar > img {
    object-fit: cover;
    border: 4px solid #fff;
    border-radius: 45px;
  }

  .ant-avatar {
    left: 10px;
  }
`

const NameContainer = styled.div`
  display: flex;
  margin-left: 2rem;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  h1 {
    margin: 0;
    width: 100%;
    text-align: left;
    font-weight: 700;
  }
  h2 {
    margin: 0;
    text-align: left;
    font-weight: 400;
  }
`

const Container = styled.div`
  min-height: 100%;
  min-width: 100%;
  display: flex;
  justify-content: center;
  padding: 10rem;
  align-items: center;
  flex-direction: column;
  @media only screen and (max-width: 900px) {
    padding: 0;
  }
`

export default withAuth(Users)
