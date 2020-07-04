import Head from "next/head"
import React from "react"
import styled from "styled-components"
import { THEME_VARIABLES } from "../config/env"

const Home = props => (
  <>
    <Head>
      <title>Aumo</title>
      <link rel="icon" href="/favicon.ico" />
    </Head>
    <Container>
      <LeftContainer>
        <Title>The receipts of the future.</Title>
        <BadgeContainer>
          <Badge src="app-store.png" style={{ opacity: 0.50, maxWidth: '33%' }} />
          <a href="https://deliprods.fra1.digitaloceanspaces.com/aumo-android-binaries/aumo.apk" style={{ maxWidth: '33%' }} >
            <Badge src="google-play.png" style={{ width: '100%' }} />
          </a>
        </BadgeContainer>
      </LeftContainer>
      <RightContainer>
        <Phone src="iphone.png" alt="iphone screenshot" />
      </RightContainer>
    </Container>
  </>
)

const Badge = styled.img`
  margin: 2.5px;
`

const BadgeContainer = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
`

const Phone = styled.img`
  height: 45%;
  width: auto;
`

const Title = styled.h1`
  font-size: 3rem;
  color: ${THEME_VARIABLES["@primary-color"]};
  text-align: center;
`

const LeftContainer = styled.div`
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 2rem;
`

const RightContainer = styled.div`
  height: 100%;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
`

const Container = styled.div`
  height: 100%;
  width: 100%;
  display: flex;
  @media only screen and (max-width: 600px) {
    margin-top: 300px;
    flex-direction: column;
  }
`

export default Home
