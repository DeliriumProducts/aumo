import Head from "next/head"
import withAuth from "../hocs/withAuth"
import styled from "styled-components"
import { Card as c, Button, Icon, message } from "antd"
import { useState, useContext } from "react"
import ModalForm from "../components/ModalForm"
import { ProductAPI } from "aumo-api"
import { BACKEND_URL } from "../config"
import { Context } from "../context/context"

export const Products = () => {
  const ctx = useContext(Context)
  const [curProduct, setCurProduct] = useState(null)
  const [loading, setLoading] = useState(true)
  const [visible, setVisible] = useState(false)
  const [formRef, setFormRef] = useState(null)

  React.useEffect(() => {
    ;(async () => {
      const data = await new ProductAPI(BACKEND_URL).getAll()
      ctx.dispatch({ type: "setProducts", payload: data })
      setLoading(false)
    })()
  }, [])

  const handleEdit = p => {
    setCurProduct(p)
    showModal()
  }

  const showModal = () => setVisible(true)

  const handleCancel = () => setVisible(false)

  const handleSubmit = () => {
    const { form } = formRef.props

    form.validateFields(async (err, product) => {
      if (err) {
        return
      }

      try {
        await new ProductAPI(BACKEND_URL).edit(curProduct.id, {
          ...product,
          price: Number(product.price),
          stock: Number(product.stock)
        })
        message.success(`Successfully edited product ${product.name}!`)
        const prods = ctx.state.products.map(pp => {
          if (pp.id === curProduct.id) {
            return {
              id: curProduct.id,
              ...product,
              stock: Number(product.stock),
              price: Number(product.price)
            }
          }
          return pp
        })
        ctx.dispatch({ type: "setProducts", payload: prods })
      } catch (err) {
        if (!err.response) {
          message.error(`${err}`, 5)
          return
        }
        if (err.response.status === 401) {
          message.error("Invalid credentials. Try again.", 1)
        } else {
          message.error("Server error, please try again")
        }
        return
      }
      form.resetFields()
      setVisible(false)
    })
  }

  const saveFormRef = fr => {
    setFormRef(fr)
  }

  return (
    <>
      <Head>
        <title>Aumo</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Container>
        {loading && ctx.state.products.length < 1 && (
          <Icon type="loading" style={{ fontSize: 24 }} spin />
        )}

        {ctx.state.products &&
          ctx.state.products.length > 0 &&
          ctx.state.products.map(p => (
            <ProductCard
              key={p.id}
              hoverable
              cover={<img alt="Product" src={p.image} />}
            >
              <StyledMeta title={p.name} description={<p>{p.description}</p>} />
              <span className="actions">
                <span>
                  <span className="price">{p.price} </span>pts.
                </span>
                <span className="actions-buttons">
                  <Button
                    size="small"
                    type="primary"
                    className="edit-button"
                    icon="edit"
                    onClick={() => handleEdit(p)}
                  ></Button>
                  <Button size="small" type="danger" icon="delete"></Button>
                </span>
              </span>
            </ProductCard>
          ))}
        <ModalForm
          wrappedComponentRef={saveFormRef}
          visible={visible}
          onCancel={handleCancel}
          onCreate={handleSubmit}
          product={curProduct}
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
    border-radius: 7px 7px 0 0;
    object-fit: cover;
    /* width: 100%; */
    height: 10rem;
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
    font-size: 15px;
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

export default withAuth(Products)
