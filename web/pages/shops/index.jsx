import { Button, Card as c, Icon, message, Popconfirm, Tooltip } from "antd"
import aumo from "aumo"
import Head from "next/head"
import Router from "next/router"
import React, { useContext, useState } from "react"
import styled from "styled-components"
import ModalForm from "../../components/ModalForm"
import ModalFormSO from "../../components/ModalFormSO"
import { Context } from "../../context/context"
import { actions } from "../../context/providers/contextProvider"
import withAuth from "../../hocs/withAuth"

export const Shops = () => {
  const ctx = useContext(Context)
  const [curShop, setCurShop] = useState(null)
  const [loading, setLoading] = useState(true)
  const [visible, setVisible] = useState(false)
  const [visibleSO, setVisibleSO] = useState(false)
  const [formRef, setFormRef] = useState(null)
  const [formRefSO, setFormRefSO] = useState(null)

  React.useEffect(() => {
    ;(async () => {
      const data = await aumo.auth.me()
      ctx.dispatch({ type: actions.SET_USER, payload: data })
      setLoading(false)
    })()
  }, [])

  const handleEdit = s => {
    setCurShop(s)
    showModal()
  }

  const showModal = () => setVisible(true)

  const handleCancel = async () => {
    setVisible(false)
    setVisibleSO(false)
    setLoading(true)

    const data = await aumo.auth.me()
    ctx.dispatch({ type: actions.SET_USER, payload: data })

    setLoading(false)
  }

  const handleSubmit = () => {
    const { form } = formRef.props

    form.validateFields(async (err, shop) => {
      if (err) {
        return
      }

      try {
        await aumo.shop.editShop(curShop.id, shop)
        message.success(`Successfully edited shop ${shop.name}! 🎉`)

        const data = await aumo.auth.me()
        ctx.dispatch({ type: actions.SET_USER, payload: data })
      } catch (err) {
        if (!err.response) {
          message.error(`${err}`, 5)
          return
        }
        if (err.response.status === 401) {
          message.error("Unauthorized.", 1)
        } else {
          message.error("Server error, please try again")
        }
        return
      }
      form.resetFields()
      setVisible(false)
    })
  }

  const handleDelete = async s => {
    try {
      await aumo.shop.deleteShop(s.id)
      message.success(`Successfully added new owner! 🎉`)

      const data = await aumo.auth.me()
      ctx.dispatch({ type: actions.SET_USER, payload: data })
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

  const saveFormRef = fr => {
    setFormRef(fr)
  }

  const saveFormRefSO = fr => {
    setFormRefSO(fr)
  }

  return (
    <>
      <Head>
        <title>Aumo</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Container>
        {loading && ctx.state.user?.shops?.length < 1 && (
          <Icon type="loading" style={{ fontSize: 24 }} spin />
        )}
        {ctx.state?.user?.shops?.length > 0 &&
          ctx.state?.user?.shops?.map(s => (
            <ProductCard
              key={s.id}
              hoverable
              cover={<img alt="Product" src={s.image} />}
              onClick={e => {
                Router.push(`/shops/[shop_id]`, `/shops/${s.id}`)
              }}
            >
              <StyledMeta title={s.name} description={<p>{s.description}</p>} />
              <span className="actions">
                <span className="actions-buttons">
                  <Tooltip placement="bottom" title="Manage owners!">
                    <Button
                      size="small"
                      type="primary"
                      icon="user"
                      onClick={e => {
                        setCurShop(s)
                        setVisibleSO(true)
                        e.stopPropagation()
                      }}
                    ></Button>
                  </Tooltip>
                  <Tooltip placement="bottom" title="Edit this shop!">
                    <Button
                      size="small"
                      type="primary"
                      className="edit-button"
                      icon="edit"
                      onClick={e => {
                        e.stopPropagation()
                        handleEdit(s)
                      }}
                    ></Button>
                  </Tooltip>
                  <Tooltip placement="bottom" title="Delete this shop!">
                    <Popconfirm
                      onConfirm={e => {
                        e.stopPropagation()
                        handleDelete(s)
                      }}
                      title={`Are you sure?`}
                      placement="bottom"
                      okText="Yes"
                      okType="danger"
                      onCancel={e => e.stopPropagation()}
                    >
                      <Button
                        size="small"
                        type="danger"
                        icon="delete"
                        onClick={e => {
                          e.stopPropagation()
                        }}
                      ></Button>
                    </Popconfirm>
                  </Tooltip>
                </span>
              </span>
            </ProductCard>
          ))}
        <ModalForm
          wrappedComponentRef={saveFormRef}
          visible={visible}
          onCancel={handleCancel}
          onCreate={handleSubmit}
          entity={curShop}
        />
        <ModalFormSO
          wrappedComponentRef={saveFormRefSO}
          visible={visibleSO}
          onCancel={handleCancel}
          shop={curShop}
          currentUser={ctx.state.user}
          cancelText="Ok"
        />
      </Container>
    </>
  )
}

const ProductCard = styled(c)`
  border-radius: 24px;
  display: flex;
  border: none;
  flex-direction: column;
  text-align: center;
  margin: 8px;
  box-shadow: rgba(0, 0, 0, 0.31) 0px 20px 24px -18px;
  width: 220px;
  height: 320px;

  .ant-card-body {
    padding-top: 0;
    height: 100%;
    text-align: left;
    & .actions {
      display: flex;
      flex-direction: row-reverse;
      width: 100%;
      justify-content: space-between;

      button {
        top: 3px;
        border: none;
        margin-left: 5px;
        border-radius: 11px;
        width: 40px;
        height: 40px;
        font-size: 18p55c353x;
        box-shadow: rgba(0, 0, 0, 0.31) 0px 20px 24px -18px;
      }

      .edit-button {
        background-color: #55c353;
      }
    }
  }

  & img {
    margin-top: 5px;
    border-radius: 7px 7px 0 0;
    object-fit: contain;
    height: 10rem;
    overflow: hidden;
  }

  .ant-card-actions {
    background-color: #fff;
    border-radius: 0 0 7px 7px;
  }

  .price {
    color: black;
    font-weight: 500;
  }
`

const StyledMeta = styled(c.Meta)`
  display: flex;
  flex-grow: 1;
  align-items: center;
  justify-content: center;
  height: 100%;
  & * {
    white-space: initial;
    overflow-wrap: normal;
  }
  .ant-card-meta-title {
    font-weight: bold;
    font-size: 20px;
    text-align: left;
  }

  .ant-card-meta-description {
    color: black;
    & p {
      /* width: 100%; */
      font-size: 14px;
      max-height: 30px;
      overflow: hidden;
      white-space: nowrap;
      text-overflow: ellipsis;
    }
  }
`

const Container = styled.div`
  display: flex;
  width: 100%;
  flex-wrap: wrap;
  justify-content: center;
  & > div {
    margin: 2rem;
  }
  @media only screen and (max-width: 600px) {
    align-items: center;
    flex-direction: column;
  }

  min-height: 100%;
  min-width: 100%;
  padding: 10rem;
  @media only screen and (max-width: 900px) {
    padding-right: 0;
    padding-bottom: 0;
    padding-left: 0;
  }
`

export default withAuth(Shops)
