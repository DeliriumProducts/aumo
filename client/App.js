import { mapping } from "@eva-design/eva"
import { NavigationContainer } from "@react-navigation/native"
import { ApplicationProvider, IconRegistry } from "@ui-kitten/components"
import { EvaIconsPack } from "@ui-kitten/eva-icons"
import aumo from "aumo"
import React from "react"
import { BACKEND_URL } from "./config"
import customM from "./mapping"
import AppNavigator from "./navigation/main"
import theme from "./theme"

if (__DEV__) {
  aumo.config.config({
    Backend: BACKEND_URL
  })
}

const App = () => {
  const [user, setUser] = React.useState(null)
  React.useEffect(() => {
    ;(async () => {
      try {
        const val = await aumo.auth.me()
        setUser(val)
      } catch (e) {}
    })()
  }, [])

  console.log(user)
  return (
    <>
      <IconRegistry icons={EvaIconsPack} />
      <ApplicationProvider
        mapping={mapping}
        theme={theme}
        customMapping={customM}
      >
        <NavigationContainer>
          <AppNavigator isAuthenticated={user != null} />
        </NavigationContainer>
      </ApplicationProvider>
    </>
  )
}

export default App
