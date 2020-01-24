import Head from "next/head"
import { Icon, Card, Button, Avatar, Modal, Carousel } from "antd"
import React from "react"
import { UserAPI } from "aumo-api"
import { THEME_VARIABLES } from "../config"
import styled from "styled-components"
import withAuth from "../hocs/withAuth.js"
import { BACKEND_URL } from "../config"

function timeout(ms) {
  return new Promise(resolve => setTimeout(resolve, ms))
}

const Users = () => {
  const [users, setUsers] = React.useState([])
  const [loading, setLoading] = React.useState(true)
  const [userModal, setUserModal] = React.useState(false)
  const [currentUser, setCurrentUser] = React.useState(null)

  React.useEffect(() => {
    ;(async () => {
      const data = await new UserAPI(BACKEND_URL).getAll()
      setUsers(data)
      setLoading(false)
    })()
  }, [])

  const showUser = async (e, user) => {
    setLoading(true)
    setUserModal(true)
    const newUser = await new UserAPI(BACKEND_URL).get(user.id)
    setLoading(false)
    setCurrentUser(newUser)
    // Modal.info({
    //   title: user.name,
    //   content: (
    //     <div>
    //       <p>some messages...some messages...</p>
    //       <p>some messages...some messages...</p>
    //     </div>
    //   ),
    //   onOk() {}
    // })
  }

  const deleteUser = (e, user) => {
    console.log(id)
  }

  return (
    <>
      <Head>
        <title>Aumo Users</title>
      </Head>
      <Container>
        {loading && users.length < 1 && (
          <Icon type="loading" style={{ fontSize: 24 }} spin />
        )}
        {users &&
          users.length > 0 &&
          users.map(u => (
            <UserCard
              key={u.id}
              id={u.id}
              user={u}
              onClick={showUser}
              onDelete={deleteUser}
            />
          ))}
        <Modal
          visible={userModal}
          centered
          onCancel={() => setUserModal(false)}
          footer={null}
        >
          <User loading={loading} user={currentUser} />
        </Modal>
      </Container>
    </>
  )
}

const User = ({ user, loading }) => {
  return (
    <>
      <Center>
        {loading && <Icon type="loading" style={{ fontSize: 24 }} spin />}
        {user && (
          <>
            <Card
              bordered={false}
              cover={
                <img
                  alt={user.email}
                  src={user.avatar}
                  style={{ width: 300 }}
                />
              }
            >
              <Card.Meta title={user.name} description={user.email} />
            </Card>
            <UserInfo>
              <div>
                <Bold>{user.points}</Bold> pts.
              </div>
              <div>
                <Bold>{user.receipts.length}</Bold> receipts
              </div>
              <div>
                <Bold>{user.orders.length}</Bold> orders
              </div>
              <Filler />
            </UserInfo>
          </>
        )}
      </Center>

      {user && user.orders && user.orders.length > 0 && (
        <Orders
          autoplay
          effect="fade"
          prevArrow={<Icon type="left-circle" style={{ color: "black" }} />}
          nextArrow={<Icon type="right-circle" style={{ color: "black" }} />}
        >
          {user.orders.map(o => (
            <Order product={o.product} key={o.order_id} />
          ))}
        </Orders>
      )}
    </>
  )
}

const Order = ({ product }) => (
  <Card
    size="small"
    height={400}
    cover={
      <img
        src={product.image}
        alt={product.name}
        height={300}
        style={{ objectFit: "cover" }}
      />
    }
  >
    <Card.Meta
      title={<span>{product.name}</span>}
      description={
        <>
          <Bold>{product.price}</Bold> pts.
        </>
      }
    />
  </Card>
)

const Orders = styled(Carousel)`
  display: flex;
  justify-content: center;
  align-items: center;
  .slick-slide {
    overflow: hidden;
    height: 400px;
  }

  .slick-dots li button {
    background: #444;
  }
  .slick-dots li.slick-active button {
    background: ${THEME_VARIABLES["@primary-color"]};
  }
`

const UserInfo = styled.div`
  display: flex;
  padding: 1rem;
  flex-direction: column;
  justify-content: center;
  align-items: center;
`

const Bold = styled.span`
  color: black;
  font-weight: 500;
`

const Center = styled.div`
  display: flex;
`

const UserCard = ({ user, onDelete, onClick }) => {
  return (
    <UserCardContainer
      hoverable
      onClick={e => {
        onClick(e, user)
      }}
    >
      <div>
        <Avatar src={user.avatar} size={80} key={user.id} className="avatar" />
      </div>
      <NameContainer>
        <h1>{user.name}</h1>
        <h2>{user.email}</h2>
      </NameContainer>
      <Filler />
      <Button
        type="danger"
        icon="delete"
        size="large"
        onClick={e => {
          e.stopPropagation()
          onDelete(e, user)
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
  height: 100%;
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
    padding-right: 0;
    padding-bottom: 0;
    padding-left: 0;
  }
`

export default withAuth(Users)
