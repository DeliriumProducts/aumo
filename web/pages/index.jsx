import React from "react"
import Head from "next/head"
import styled from "styled-components"

const Home = () => (
  <>
    <Head>
      <title>Aumo</title>
      <link rel="icon" href="/favicon.ico" />
    </Head>
    <Container>
      <LeftContainer>
        <Title>The receipts of the future.</Title>
        <BadgeContainer>
          <Badge src="app-store.png" />
          <Badge src="google-play.png" />
        </BadgeContainer>
      </LeftContainer>
      <RightContainer>
        <Phone src="iphone.png" alt="iphone screenshot" />
      </RightContainer>
    </Container>
  </>
)

const Badge = styled.img`
  width: 20%;
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
  color: #083aa4;
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
`

export default Home
