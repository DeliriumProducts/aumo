import React from "react"
import {
  ApplicationProvider,
  Layout,
  Text,
  IconRegistry,
  Button,
  Icon
} from "@ui-kitten/components"
import { mapping } from "@eva-design/eva"
import Container from "./navigation/Container"
import { EvaIconsPack } from "@ui-kitten/eva-icons"
import theme from "./theme"
import customM from "./mapping"

export const FacebookIcon = style => <Icon name="facebook" {...style} />

export const LoginButton = () => (
  <Button icon={FacebookIcon}>Login with Facebook</Button>
)

const App = () => (
  <>
    <IconRegistry icons={EvaIconsPack} />
    <ApplicationProvider
      mapping={mapping}
      theme={theme}
      customMapping={customM}
    >
      <Container />
    </ApplicationProvider>
  </>
)

export default App
