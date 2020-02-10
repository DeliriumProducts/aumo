import { Button, Spinner } from "@ui-kitten/components"
import aumo from "aumo"
import React from "react"
import { View } from "react-native"
import { Context } from "../../context/context"
import { actions } from "../../context/providers/provider"

export default () => {
  const ctx = React.useContext(Context)
  const [loading, setLoading] = React.useState(false)
  return (
    <>
      <Button
        disabled={loading}
        size="large"
        appearance="ghost"
        onPress={async () => {
          try {
            setLoading(true)
            await aumo.auth.logout()
            ctx.dispatch({ type: actions.SET_USER, payload: null })
          } catch (error) {
          } finally {
            setLoading(false)
          }
        }}
      >
        Logout!
      </Button>
      <View
        style={{
          justifyContent: "center",
          alignItems: "center",
          height: "100%"
        }}
      >
        {loading && <Spinner size="giant" />}
      </View>
    </>
  )
}
