import App from "next/app"
import Nav from "../components/Nav"
import Router from "next/router"
import { Affix } from "antd"
import NProgress from "nprogress"
import React from "react"
import styled, { createGlobalStyle } from "styled-components"
import "../assets/nprogress.less"
import ContextProvider from "../context/providers/contextProvider"
import { Context } from "../context/context.js"

/**
 * https://github.com/zeit/next.js/tree/canary/examples/with-loading
 */

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
       background-color: #f9faff;
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

const BlobWrapper = styled.div`
  min-height: 100vh;
  position: relative;
  overflow: hidden;
`

const Blob = styled.img`
  position: absolute;
  z-index: 0;
  @media only screen and (max-width: 900px) {
    display: none;
  }
`

const Layout = ({ children, route }) => {
  const ctx = React.useContext(Context)
  console.log(ctx.user)

  return (
    <>
      <BlobWrapper>
        <Blob
          src="blob-shape.svg"
          style={{
            top: "-8vmax",
            right: "-10vmax",
            transform: "rotate(90deg)"
          }}
        />
        <Blob
          src="blob-shape-2.svg"
          style={{
            left: "-8vmax",
            bottom: "-13vmax"
          }}
        />
        <Blob
          src="blob-shape-3.svg"
          style={{
            bottom: "-13vmax",
            right: "-7vmax"
          }}
        />
      </BlobWrapper>
      <Affix offsetTop={0}>
        <Nav name={ctx.state.user?.name} route={route} />
      </Affix>
      {children}
    </>
  )
}
