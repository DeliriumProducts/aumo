import aumo from "aumo"
import React from "react"
import { Button } from "react-native"
import { Context } from "../../context/context"
import { actions } from "../../context/providers/provider"

export default () => {
  const ctx = React.useContext(Context)
  return (
    <Button
      title="Logout!"
      onPress={async () => {
        await aumo.auth.logout()
        ctx.dispatch({ type: actions.SET_USER, payload: null })
      }}
    ></Button>
  )
}
