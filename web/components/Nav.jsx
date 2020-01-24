import React from "react"
import Router from "next/router"
import { BACKEND_URL } from "../config"
import { AuthAPI } from "aumo-api"
import Link from "next/link"
import styled from "styled-components"
import { Icon, Divider, Button, message } from "antd"

const links = [
  { href: "/products", label: "Products", icon: <Icon type="shop" /> },
  { href: "/users", label: "Users", icon: <Icon type="user" /> }
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
          <Button type="primary">GO TO ADMIN PANEL</Button>
        </Link>
      ) : props.route === "/login" ? (
        <></>
      ) : (
        <>
          <Welcome>
            Welcome back, <span>{props.name}</span>
          </Welcome>
          {props.route === "/products" ? (
            <Button type="primary" icon="plus" className="new-button">
              NEW
            </Button>
          ) : (
            <></>
          )}
          <LinkList>
            {links.map(({ key, href, label, icon }) => (
              <Link key={key} href={href}>
                <LinkItem isSelected={props.route === href}>
                  {icon}
                  {label}
                </LinkItem>
              </Link>
            ))}
            <Divider type="vertical" />
            <Button
              type="ghost"
              onClick={async () => {
                await new AuthAPI(BACKEND_URL).logout()
                message.success("Logged out!")
                Router.replace("/")
              }}
            >
              <Icon type="logout" />
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
  box-shadow: rgba(0, 0, 0, 0.31) 0px 20px 24px -24px;
  align-items: center;
  z-index: 999;
  @media only screen and (max-width: 600px) {
    align-items: center;
    justify-content: center;
    flex-direction: column;
  }

  .new-button {
    background-color: #55c353;
    border: none;
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
    width: 100%;
    margin-top: 1%;
    align-items: center;
    flex-direction: row;
    padding-left: 0;
    justify-content: center;
  }
`

const LinkItem = styled.a`
  font-family: "Montserrat";
  color: ${props => (props.isSelected ? "#fff" : "#083aa4")};
  background-color: ${props => (props.isSelected ? "#083aa4" : "")};
  font-size: 1rem;
  font-weight: 500;
  display: flex;
  justify-content: center;
  align-items: center;
  margin-left: 5px;
  text-decoration: none;
  border-radius: 10px;
  padding: 8px;
  i {
    margin-right: 5px;
  }
`

export default Nav
