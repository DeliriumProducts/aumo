import { mapping } from "@eva-design/eva"
import { NavigationContainer } from "@react-navigation/native"
import { ApplicationProvider, IconRegistry } from "@ui-kitten/components"
import { EvaIconsPack } from "@ui-kitten/eva-icons"
import aumo from "aumo"
import React from "react"
import { SafeAreaView } from "react-native"
import { BACKEND_URL } from "./config"
import { Context } from "./context/context"
import ContextProvider, { actions } from "./context/providers/provider"
import customM from "./mapping"
import AppNavigator from "./navigation/main"
import theme from "./theme"

if (__DEV__) {
  aumo.config.config({
    Backend: BACKEND_URL
  })
}

const App = () => {
  const ctx = React.useContext(Context)

  React.useEffect(() => {
    ;(async () => {
      try {
        const val = await aumo.auth.me()
        ctx.dispatch({ type: actions.SET_USER, payload: val })
      } catch (e) {}
    })()
  }, [])

  return (
    <>
      <IconRegistry icons={EvaIconsPack} />
      <ApplicationProvider
        mapping={mapping}
        theme={theme}
        customMapping={customM}
      >
        <SafeAreaView style={{ flex: 1 }}>
          <NavigationContainer>
            <AppNavigator isAuthenticated={ctx.state.user != null} />
          </NavigationContainer>
        </SafeAreaView>
      </ApplicationProvider>
    </>
  )
}

export default () => (
  <ContextProvider>
    <App />
  </ContextProvider>
)
