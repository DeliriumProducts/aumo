import React from "react"
import Head from "next/head"
import styled from "styled-components"
import { Input, Icon, Form, Button, message } from "antd"
import { AuthAPI } from "aumo-api"
import { BACKEND_URL } from "../config"
import Router from "next/router"

const FormItem = Form.Item

const Login = props => {
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

        const authAPI = new AuthAPI(BACKEND_URL)
        try {
          await authAPI.login(credentials)
        } catch (e) {
          message.error(`${e.response.data.error}`, 5)
          return
        }

        Router.replace("/products")
      }
    })
  }
  return (
    <>
      <Head>
        <title>Aumo</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Container>
        <Card>
          <Form onSubmit={handleSubmit} className="login-form">
            <p>Login to manage aumo</p>
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
                htmlType="submit"
                className="login-form-button"
              >
                Login
              </Button>
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

  return { auth }
}

const Card = styled.div`
  background-color: #fff;
  text-align: center;
  padding: 20px;
  border-radius: 30px;
  box-shadow: 0px 2px 2px rgba(0, 0, 0, 0.12);
  display: flex;

  & p {
    color: black;
    font-weight: 600;
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
