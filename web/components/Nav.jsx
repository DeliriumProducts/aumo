import React from "react"
import Router from "next/router"
import { BACKEND_URL } from "../config"
import { AuthAPI } from "aumo-api"
import Link from "next/link"
import styled from "styled-components"
import { Divider, Button } from "antd"

const links = [
  { href: "/products", label: "Products" },
  { href: "/users", label: "Users" }
].map(link => ({
  ...link,
  key: `nav-link-${link.href}-${link.label}`
}))

const Nav = props => (
  <nav>
    <Menu>
      <Link href={"/"}>
        <Logo src="aumo.png" />
      </Link>
      {props.route === "/" ? (
        <Link href="/login">
          <Button type="primary">LOGIN NOW</Button>
        </Link>
      ) : props.route === "/login" ? (
        <></>
      ) : (
        <>
          <Welcome>
            Welcome back, <span>{props.name}</span>
          </Welcome>
          <LinkList>
            {links.map(({ key, href, label }) => (
              <Link key={key} href={href}>
                <LinkItem>{label}</LinkItem>
              </Link>
            ))}
            <Divider type="vertical" className="divider" />
            <Button
              type="primary"
              onClick={async () => {
                await new AuthAPI(BACKEND_URL).logout()
                Router.replace("/")
              }}
            >
              LOGOUT
            </Button>
          </LinkList>
        </>
      )}
    </Menu>
  </nav>
)

const Logo = styled.img`
  cursor: pointer;
  max-width: 10%;
  height: auto;
  margin-bottom: 9px;
  @media only screen and (max-width: 600px) {
    max-width: 20%;
  }
`

const Welcome = styled.div`
  width: 100%;
  color: black;
  align-self: center;
  margin-left: 4rem;
  text-align: left;
  font-family: "Montserrat";
  font-size: 1rem;
  font-weight: 500;
  text-decoration: none;
  @media only screen and (max-width: 600px) {
    text-align: center;
    margin-left: 0;
  }

  span {
    text-align: left;
    font-weight: bold;
    font-family: "Montserrat";
    color: #083aa4;
    font-size: 1rem;
    text-decoration: none;
  }
`

const Menu = styled.div`
  position: fixed;
  width: 100%;
  display: flex;
  background-color: #fff;
  padding: 1.2rem;
  justify-content: space-between;
  align-items: center;
  z-index: 999;
  @media only screen and (max-width: 600px) {
    align-items: center;
    justify-content: center;
    flex-direction: column;
    .divider {
      display: none;
    }
  }
`
const LinkList = styled.ul`
  display: flex;
  margin-bottom: 0;
  margin-top: 0;
  justify-content: center;
  align-items: center;
  flex-direction: row;
  @media only screen and (max-width: 600px) {
    align-items: center;
    padding-left: 0;
    justify-content: center;
    flex-direction: column;
  }
`

const LinkItem = styled.a`
  font-family: "Montserrat";
  color: #083aa4;
  font-size: 1.2rem;
  font-weight: 700;
  text-decoration: none;
  padding: 10px;
`

export default Nav
