import React from "react"
import {
  ApplicationProvider,
  Layout,
  Text,
  Button
} from "@ui-kitten/components"
import { mapping } from "@eva-design/eva"
import theme from "./theme"

const HomeScreen = () => (
  <Layout style={{ flex: 1, justifyContent: "center", alignItems: "center" }}>
    <Text category="h1">HOME</Text>
    <Button>hiii!</Button>
  </Layout>
)

const App = () => (
  <ApplicationProvider mapping={mapping} theme={theme}>
    <HomeScreen />
  </ApplicationProvider>
)

export default App
