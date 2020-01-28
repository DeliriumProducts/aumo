import React from "react"
import { ApplicationProvider, IconRegistry } from "@ui-kitten/components"
import { mapping } from "@eva-design/eva"
import Container from "./navigation/container"
import { EvaIconsPack } from "@ui-kitten/eva-icons"
import theme from "./theme"
import customM from "./mapping"

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
