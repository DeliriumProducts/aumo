import React, { Component } from "react"

export const withAuth = (C, roles = []) =>
  class extends Component {
    static async getInitialProps(ctx) {
      const { req, res } = ctx
      let auth = {
        user: null,
        isAuthenticated: false
      }
      /**
       * Check wheter authentication is happening server-side or client-side based on received context
       */
      if (req && res) {
        if (req.headers.cookie) {
          auth = await StaffAPI.isAuthenticated(req.headers.cookie)
        }
        if (!auth.isAuthenticated) {
          res.writeHead(302, {
            Location: "/login"
          })
          res.end()
        } else if (
          roles.length &&
          auth.user &&
          auth.user.role !== "Admin" &&
          !roles.includes(auth.user.role)
        ) {
          res.writeHead(302, {
            Location: "/admin"
          })
          res.end()
        }
      } else {
        auth = await StaffAPI.isAuthenticated()
        if (!auth.isAuthenticated) {
          Router.replace("/login")
        } else if (
          roles.length &&
          auth.user &&
          auth.user.role !== "Admin" &&
          !roles.includes(auth.user.role)
        ) {
          Router.replace("/admin")
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
        user: auth.user
      }
    }
    componentDidMount() {
      this.context.dispatch({ type: "setUser", payload: this.props.user })
    }
    render() {
      return React.createElement(C, Object.assign({}, this.props))
    }
  }
