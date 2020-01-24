import React from "react"
import Head from "next/head"
import styled from "styled-components"

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
          <Badge src="app-store.png" />
          <Badge src="google-play.png" />
        </BadgeContainer>
        <a style={{ marginTop: "3%" }} href="paper.pdf">
          Learn more
        </a>
      </LeftContainer>
      <RightContainer>
        <Phone src="iphone.png" alt="iphone screenshot" />
      </RightContainer>
    </Container>
  </>
)

const Badge = styled.img`
  padding: 5px;
  width: 25%;
  @media only screen and (max-width: 600px) {
    width: 35%;
  }
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
  @media only screen and (max-width: 600px) {
    margin-top: 300px;
    flex-direction: column;
  }
`

export default Home
