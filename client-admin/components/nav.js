import React from "react"
import Link from "next/link"
import styled from "styled-components"
import { Divider } from "antd"

const links = [
  { href: "/products", label: "Products" },
  { href: "/users", label: "Users" }
].map(link => ({
  ...link,
  key: `nav-link-${link.href}-${link.label}`
}))

const Nav = ({ props }) => (
  <nav>
    <Menu>
      <img src="aumo.png" className="aumo-logo" />
      <div className="welcome-text">
        Welcome back, <strong className="welcome-name">Nasko</strong>
      </div>
      <LinkList>
        {links.map(({ key, href, label }) => (
          <Link key={key} href={href}>
            <LinkItem>{label}</LinkItem>
          </Link>
        ))}
        <Divider type="vertical" />
        <Link href={"/login"}>
          <LinkItem type="primary">LOGOUT</LinkItem>
        </Link>
      </LinkList>
    </Menu>

    <style jsx>{`
      https://nextjs.org
      :global(body) {
        margin: 0;
        font-family: -apple-system, BlinkMacSystemFont, Avenir Next, Avenir,
          Helvetica, sans-serif;
      }
    `}</style>
  </nav>
)

const Menu = styled.div`
  display: flex;
  padding: 0.5rem;
  justify-content: space-between;
  .aumo-logo {
    width: 10%;
    align-self: center;
    height: 20%;
  }

  .welcome-text {
    width: 100%;
    color: black;
    align-self: center;
    margin-left: 4rem;
    text-align: left;
    font-family: "Montserrat";
    font-size: 1rem;
    font-weight: 500;
    text-decoration: none;
  }

  .welcome-name {
    text-align: left;
    font-family: "Montserrat";
    color: #083aa4;
    font-size: 1rem;
    text-decoration: none;
  }

  @media only screen and (max-width: 600px) {
    align-items: center;
    justify-content: center;
    flex-direction: column;
  }
`
const LinkList = styled.ul`
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: row;
  @media only screen and (max-width: 600px) {
    align-items: center;
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
