import {
  Avatar,
  Button,
  Card,
  Carousel,
  Icon,
  message,
  Modal,
  Popconfirm,
  Radio,
  Tag
} from "antd"
import RadioGroup from "antd/lib/radio/group"
import aumo from "aumo"
import Head from "next/head"
import React from "react"
import styled from "styled-components"
import { THEME_VARIABLES } from "../config"
import { Context } from "../context/context.js"
import withAuth from "../hocs/withAuth.js"

const colors = {
  ["Shop Owner"]: "purple",
  Admin: "magenta"
}

const Users = () => {
  const ctx = React.useContext(Context)
  const [users, setUsers] = React.useState([])
  const [role, setRole] = React.useState("Admin")
  const [loading, setLoading] = React.useState(true)
  const [userModal, setUserModal] = React.useState(false)
  const [currentUser, setCurrentUser] = React.useState(null)

  React.useEffect(() => {
    ;(async () => {
      const data = await aumo.user.getAllUsers()
      setUsers(data)
      setLoading(false)
    })()
  }, [])

  const showUser = async (_, user) => {
    setCurrentUser(null)
    setLoading(true)
    setUserModal(true)
    try {
      const newUser = await aumo.user.getUser(user.id)
      setCurrentUser(newUser)
    } catch (e) {
      message.error(`${e.error}`)
      setUserModal(false)
    } finally {
      setLoading(false)
    }
  }

  const handleRoleChange = role => {
    setRole(role)
  }

  const changeRole = async (e, user) => {
    if (user.role === role) {
      return
    }
    try {
      await aumo.user.setRole(user.id, role)
      message.success(`Successfully changed ${user.name}'s role to ${role}! 🎉`)
    } catch (err) {
      if (!err.response) {
        message.error(`${err}`, 5)
        return
      }
      if (err.response.status === 401) {
        message.error("Unathorized. Try again.", 1)
      } else {
        message.error("Server error, please try again")
      }
      return
    }

    setUsers(users =>
      users.map(pu => {
        if (pu.id == user.id) {
          return {
            ...pu,
            role: role
          }
        }

        return pu
      })
    )
  }

  const deleteUser = async user => {
    try {
      await aumo.user.deleteUser(user.id)
      message.success(`Successfully deleted user ${user.name}! 🎉`)
      setUsers(prevUsers =>
        prevUsers.filter(pu => {
          return pu.id !== user.id
        })
      )
    } catch (err) {
      if (!err.response) {
        message.error(`${err}`, 5)
        return
      }
      if (err.response.status === 401) {
        message.error("Unathorized. Try again.", 1)
      } else {
        message.error("Server error, please try again")
      }
      return
    }
  }

  const addPoints = async user => {
    try {
      await aumo.user.addPoints(user.id, 500)
      message.success(`Successfully added 500 points to user ${user.name}! 🎉`)
    } catch (err) {
      if (!err.response) {
        message.error(`${err}`, 5)
        return
      }
      if (err.response.status === 401) {
        message.error("Unathorized. Try again.", 1)
      } else {
        message.error("Server error, please try again")
      }
      return
    }
  }

  const subPoints = async user => {
    try {
      await aumo.user.subPoints(user.id, 500)
      message.success(
        `Successfully removed 500 points from user ${user.name}! 🎉`
      )
    } catch (err) {
      if (!err.response) {
        message.error(`${err}`, 5)
        return
      }
      if (err.response.status === 401) {
        message.error("Unathorized. Try again.", 1)
      } else {
        message.error("Server error, please try again")
      }
      return
    }
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
              myEmail={ctx.state.user?.email}
              key={u.id}
              user={u}
              onClick={showUser}
              onDelete={deleteUser}
              handleRoleChange={handleRoleChange}
              changeRole={changeRole}
              addPoints={addPoints}
              subPoints={subPoints}
              role={role}
            />
          ))}
        <Modal
          width={400}
          visible={userModal && currentUser != null}
          centered
          onCancel={() => {
            setUserModal(false)
          }}
          footer={null}
        >
          <User loading={loading} user={currentUser} />
        </Modal>
      </Container>
    </>
  )
}

const UserCard = ({
  myEmail,
  user,
  onDelete,
  onClick,
  handleRoleChange,
  changeRole,
  role,
  addPoints,
  subPoints
}) => {
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
        <div className="role">
          <h1>{user.name}</h1>
          <Tag color={colors[user.role]}>{user.role.toUpperCase()}</Tag>
        </div>
        <h2>{user.email}</h2>
      </NameContainer>
      <Filler />
      <Button
        icon="plus"
        shape="circle"
        onClick={e => {
          e.stopPropagation()
          addPoints(user)
        }}
        style={{
          marginRight: 20
        }}
      ></Button>
      <Button
        icon="minus"
        shape="circle"
        onClick={e => {
          e.stopPropagation()
          subPoints(user)
        }}
        style={{
          marginRight: 20
        }}
      ></Button>
      <div onClick={e => e.stopPropagation()}>
        <Popconfirm
          icon={<Icon type="team" style={{ color: "unset" }} />}
          placement="bottom"
          onCancel={e => e.stopPropagation()}
          disabled={myEmail === user.email}
          onConfirm={e => {
            e.stopPropagation()
            changeRole(e, user)
          }}
          title={
            <>
              <RadioGroup
                style={{ display: "flex", flexDirection: "column" }}
                onChange={e => {
                  e.stopPropagation()
                  const role = e.target.value
                  handleRoleChange(role)
                }}
                value={role}
              >
                <span style={{ fontWeight: 500, marginBottom: 5 }}>
                  Available Roles
                </span>
                <Radio value={"Customer"}>Customer</Radio>
                <Radio value={"Admin"}> Admin</Radio>
              </RadioGroup>
            </>
          }
        >
          <Button
            icon="edit"
            disabled={myEmail === user.email}
            onClick={e => {
              e.stopPropagation()
              handleRoleChange(user.role)
            }}
            style={{
              marginRight: 20
            }}
          >
            Change role
          </Button>
        </Popconfirm>
      </div>
      <Popconfirm
        onConfirm={e => {
          e.stopPropagation()
          onDelete(user)
        }}
        disabled={myEmail === user.email}
        title={`Are you sure?`}
        placement="bottom"
        okText="Yes"
        okType="danger"
        onCancel={e => e.stopPropagation()}
      >
        <Button
          type="danger"
          icon="delete"
          disabled={myEmail === user.email}
          onClick={e => e.stopPropagation()}
          style={{
            right: 10
          }}
        >
          Delete
        </Button>
      </Popconfirm>
    </UserCardContainer>
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
              <Card.Meta
                title={
                  <>
                    {user.name}
                    <Tag
                      color={colors[user.role]}
                      key={user.role}
                      style={{ marginLeft: 10 }}
                    >
                      {user.role.toUpperCase()}
                    </Tag>
                  </>
                }
                description={user.email}
              />
              <UserInfo>
                <div>
                  <Bold>{user.points}</Bold> pts.
                </div>
                <div>
                  <Bold>{user.receipts.length}</Bold>{" "}
                  {user.receipts.length == 1 ? "receipt" : "receipts"}
                </div>
                <div>
                  <Bold>{user.orders.length}</Bold>{" "}
                  {user.orders.length == 1 ? "order" : "orders"}
                </div>
              </UserInfo>
            </Card>
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
  margin-top: 8px;
`

const Bold = styled.span`
  color: black;
  font-weight: 500;
`

const Center = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
`

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
  * h1 {
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
  .role {
    width: 100%;
  }
  .role span {
    margin-top: 5px;
    margin-bottom: 5px;
    float: left;
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
