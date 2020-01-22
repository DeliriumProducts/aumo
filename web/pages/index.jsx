import React from "react"
import Head from "next/head"
import Nav from "../components/Nav"

const Home = () => (
  <div>
    <Head>
      <title>Aumo</title>
      <link rel="icon" href="/favicon.ico" />
    </Head>
    <Nav name="Nasko" />
    <div className="hero">Welcome!</div>
  </div>
)

export default Home
