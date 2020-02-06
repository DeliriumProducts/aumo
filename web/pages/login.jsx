import { Button, Form, Icon, Input, message } from "antd"
import aumo from "aumo"
import Head from "next/head"
import Link from "next/link"
import Router from "next/router"
import React from "react"
import styled from "styled-components"
import { BACKEND_URL } from "../config"

const FormItem = Form.Item

const Login = props => {
  const [loading, setLoading] = React.useState(false)
  const { getFieldDecorator } = props.form
  const handleSubmit = e => {
    e.preventDefault()
    props.form.validateFields(async (err, values) => {
      if (!err) {
        const { email, password } = values

        const credentials = {
          email,
          password
        }

        setLoading(true)
        try {
          await aumo.auth.login(credentials)
          message.success("Logged in!", 3, () => Router.replace("/products"))
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
        } finally {
          setLoading(false)
        }
      }
    })
  }

  return (
    <>
      <Head>
        <title>Aumo Login</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Container>
        <Card>
          <Form onSubmit={handleSubmit} className="login-form">
            <p>Login to manage Aumo</p>
            <FormItem>
              {getFieldDecorator("email", {
                rules: [
                  {
                    type: "email",
                    message: "The input is not valid Email!"
                  },
                  {
                    required: true,
                    message: "Please input your Email!"
                  }
                ]
              })(
                <Input
                  prefix={
                    <Icon type="mail" style={{ color: "rgba(0,0,0,.25)" }} />
                  }
                  type="email"
                  placeholder="Email"
                />
              )}
            </FormItem>
            <FormItem>
              {getFieldDecorator("password", {
                rules: [
                  {
                    required: true,
                    message: "Please input your password!"
                  }
                ]
              })(
                <Input.Password
                  prefix={
                    <Icon type="lock" style={{ color: "rgba(0,0,0,.25)" }} />
                  }
                  type="password"
                  placeholder="Password"
                />
              )}
            </FormItem>
            <FormItem>
              <Button
                type="primary"
                loading={loading}
                htmlType="submit"
                className="login-form-button"
              >
                Login
              </Button>
              Or{" "}
              <Link href="/register">
                <a>register now!</a>
              </Link>
            </FormItem>
          </Form>
        </Card>
      </Container>
    </>
  )
}

Login.getInitialProps = async ctx => {
  const { req, res } = ctx
  let auth = {}
  /**
   * Check wheter authentication is happening server-side or client-side based on received context
   */
  if (req && res) {
    if (req.headers.cookie) {
      try {
        auth = await new AuthAPI(BACKEND_URL).me(req.headers.cookie)
        if (auth.role === "Admin") {
          res.writeHead(302, {
            Location: "/products"
          })
          res.end()
        }
      } catch (err) {}
    }
  } else {
    try {
      auth = await new AuthAPI(BACKEND_URL).me()
      if (auth.role === "Admin") {
        Router.replace("/products")
      }
    } catch (err) {}
  }

  return { user: !!auth }
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

  .login-form-button {
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

const WrappedLogin = Form.create()(Login)

export default WrappedLogin
