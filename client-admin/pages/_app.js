import App from "next/app"
import React from "react"

const GlobalStyle = createGlobalStyle`
  html,
  body {
    margin: 0;
    height: auto;
    min-height: 100%;
    scroll-behavior: smooth;
  }

  body {
    min-height: 100%;
    height: initial;
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
    height: 100%;
    min-height: 100%;
  }
`

export default class MyApp extends App {
  render() {
    const { Component, pageProps } = this.props
    return (
      <>
        <GlobalStyle />
        <Component {...pageProps} />
      </>
    )
  }
}
