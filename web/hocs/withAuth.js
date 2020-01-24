import React, { Component } from "react"
import Router from "next/router"
import { AuthAPI } from "aumo-api"
import { BACKEND_URL } from "../config"
import { Context } from "../context/context.js"

export default C =>
  class extends Component {
    static contextType = Context

    static async getInitialProps(ctx) {
      const { req, res } = ctx
      let auth = {}

      /**
       * Check wheter authentication is happening server-side or client-side based on received context
       */
      if (req && res) {
        if (req.headers.cookie) {
          try {
            auth = await new AuthAPI(BACKEND_URL).me(req.headers.cookie)
            if (auth.role !== "Admin") {
              throw {
                status: 401
              }
            }
          } catch (err) {
            if (err.status === 401) {
              res.writeHead(302, {
                Location: "/login"
              })
              res.end()
            }
          }
        } else {
          res.writeHead(302, {
            Location: "/login"
          })
          res.end()
        }
      } else {
        try {
          auth = await new AuthAPI(BACKEND_URL).me()
          if (auth.role !== "Admin") {
            throw {
              status: 401
            }
          }
        } catch (err) {
          if (err.status === 401) {
            Router.replace("/login")
          }
        }
      }

      /**
       * Call the getInitalProps of the wrapped component
       */
      const composedInitialProps = C.getInitialProps
        ? await C.getInitialProps(ctx)
        : {}
      return {
        ...composedInitialProps,
        user: auth
      }
    }

    componentDidMount() {
      this.context.dispatch({ type: "setUser", payload: this.props.user })
    }
    render() {
      return <C {...this.props} />
    }
  }
