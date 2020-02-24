import { Button, Form, Icon, Input, message } from "antd"
import aumo from "aumo"
import Head from "next/head"
import styled from "styled-components"
import withAuth from "../hocs/withAuth"

const LOL = () => {
  const [totalSum, setTotalSum] = React.useState(0)
  const [content, setContent] = React.useState("")
  const [shopID, setShopID] = React.useState(null)

  return (
    <>
      <Head>
        <title>Receipts generator</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Container>
        <Card>
          <Form
            onSubmit={async e => {
              e.preventDefault()
              try {
                const receipt = {
                  shop_id: Number(shopID),
                  content: content,
                  total: Number(totalSum)
                }

                const response = await aumo.receipt.createReceipt(receipt)
                message.success(response.receipt_id, 10)
              } catch (error) {
                message.error(error, 10)
              }
            }}
            className="register-form"
          >
            <Input
              prefix={<Icon type="shop" style={{ color: "rgba(0,0,0,.25)" }} />}
              onChange={e => {
                let a = e.target.value
                setShopID(a)
              }}
              placeholder="Shop ID"
            />
            <Input.TextArea
              style={{ marginTop: 10 }}
              prefix={
                <Icon type="file-text" style={{ color: "rgba(0,0,0,.25)" }} />
              }
              rows={4}
              onChange={e => {
                let a = e.target.value
                setContent(a)
              }}
              placeholder="Content"
            />
            <Input
              style={{ marginTop: 10 }}
              prefix={
                <Icon type="dollar" style={{ color: "rgba(0,0,0,.25)" }} />
              }
              onChange={e => {
                let a = e.target.value
                setTotalSum(a)
              }}
              placeholder="Total"
            />
            <Button
              style={{ marginTop: 10 }}
              type="primary"
              htmlType="submit"
              className="register-form-button"
            >
              Generate
            </Button>
          </Form>
        </Card>
      </Container>
    </>
  )
}

const Card = styled.div`
  background-color: #fff;
  text-align: center;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 2rem;
  border-radius: 30px;
  box-shadow: rgba(0, 0, 0, 0.31) 0px 20px 24px -18px;
  display: flex;

  & p {
    color: black;
    font-weight: 600;
  }

  .register-form {
    max-width: 300px;
  }

  .register-form-button {
    width: 100%;
  }
`

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

export default withAuth(LOL)
