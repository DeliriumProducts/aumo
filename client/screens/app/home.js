import { Button, Text } from "@ui-kitten/components"
import React from "react"
import { View } from "react-native"
import NfcManager, { NdefParser, NfcEvents } from "react-native-nfc-manager"

export default () => {
  React.useEffect(() => {
    NfcManager.start()
    NfcManager.setEventListener(NfcEvents.DiscoverTag, tag => {
      let msgs = tag.ndefMessage.map(NdefParser.parseText)
      console.warn(msgs)
    })

    return () => {
      NfcManager.setEventListener(NfcEvents.DiscoverTag, null)
      NfcManager.unregisterTagEvent().catch(() => 0)
    }
  }, [])

  return (
    <View>
      <Text>aumo!</Text>
      <Button
        onPress={async () => {
          try {
            await NfcManager.registerTagEvent()
          } catch (error) {
            console.warn(error)
            NfcManager.unregisterTagEvent().catch(() => 0)
          }
        }}
      >
        Click to scan!
      </Button>
    </View>
  )
}
