import App from "next/app"
import Nav from "../components/Nav"
import Router from "next/router"
import { Affix } from "antd"
import NProgress from "nprogress"
import React from "react"
import { createGlobalStyle } from "styled-components"
import "../assets/nprogress.less"

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
    height: auto;
    min-height: 100%;
    scroll-behavior: smooth;
  }

  body {
    background-color: #F9FAFF;
    /* background-image: url('/static/circles-primary.svg'),
    url('/static/circles-accent.svg');
    background-repeat: repeat-y;
    background-size: 85vmax, 65vmax;
    background-position: bottom -20vmax left -60vmax, top -10vmax right -36vmax;
    @media screen and (max-width: 768px) {
      background-position: bottom -10vmax left -60vmax, top -80vmax right -36vmax;
      background-size: 80vmax, 100%;
    } */
  }

  * {
    box-sizing: border-box;
    font-family: Montserrat, 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  }

    #__next,
    #__next > div {
        min-height: 100vh;
        min-width: 100vw;
        position: absolute;
  }
`

export default class MyApp extends App {
  render() {
    const { Component, pageProps } = this.props
    return (
      <>
        <Affix offsetTop={0}>
          <Nav name="Nasko" />
        </Affix>
        <GlobalStyle />
        <Component {...pageProps} />
      </>
    )
  }
}
