import App from "next/app"
import Nav from "../components/Nav"
import Router from "next/router"
import NProgress from "nprogress"
import React from "react"
import styled, { createGlobalStyle } from "styled-components"
import "../assets/nprogress.less"
import ContextProvider from "../context/providers/contextProvider"
import { Context } from "../context/context.js"

Router.events.on("routeChangeStart", () => {
  NProgress.start()
})
Router.events.on("routeChangeComplete", () => NProgress.done())
Router.events.on("routeChangeError", () => NProgress.done())

const GlobalStyle = createGlobalStyle`
  html,
  body {
    margin: 0;
    padding: 0;
    height: auto;
    min-height: 100%;
    scroll-behavior: smooth;
    box-sizing: border-box
  }

   body {
    min-height: 100%;
    height: initial;
    background-image: url('blob-shape.svg'), url('blob-shape-2.svg');
    background-repeat: no-repeat;
    background-attachment: fixed;
    background-position: 
      top -8vmax
      right -10vmax,
      left -8vmax
      bottom -13vmax;

      @media only screen and (max-width: 1366px) {
            background-position: 
              top -8vmax
              right -20vmax,
              left -8vmax
              bottom -34vmax;
      }

      @media only screen and (max-width: 600px) {
          background: #F9FAFF;
      }

   }

  * {
    box-sizing: border-box;
    font-family: Montserrat, 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  }

    #__next,
    #__next > div {
        min-height: 100vh;
        min-width: 100%;
        position: absolute;
        box-sizing: border-box
  }
`

export default class MyApp extends App {
  render() {
    const {
      Component,
      pageProps,
      router: { route }
    } = this.props

    return (
      <ContextProvider>
        <GlobalStyle />
        <Layout route={route}>
          <Component {...pageProps} />
        </Layout>
      </ContextProvider>
    )
  }
}

const Layout = ({ children, route }) => {
  const ctx = React.useContext(Context)

  return (
    <>
      <Nav name={ctx.state.user?.name} route={route} />
      {children}
    </>
  )
}
