import { Icon, Text } from "@ui-kitten/components"
import React from "react"
import { Image, View } from "react-native"
import { TouchableOpacity } from "react-native-gesture-handler"
import { NeomorphBox } from "react-native-neomorph-shadows"
import NfcManager, { NdefParser, NfcEvents } from "react-native-nfc-manager"
import theme from "../../theme"

export default () => {
  React.useEffect(() => {
    NfcManager.start()
    NfcManager.setEventListener(NfcEvents.DiscoverTag, tag => {
      console.log(tag)
      let msgs = tag.ndefMessage.map(NdefParser.parseText)
      console.warn(msgs)
    })

    return () => {
      NfcManager.setEventListener(NfcEvents.DiscoverTag, null)
      NfcManager.unregisterTagEvent().catch(() => 0)
    }
  }, [])

  return (
    <View
      style={{
        justifyContent: "space-between",
        alignItems: "center",
        height: "100%"
      }}
    >
      <Image
        source={require("../../assets/AumoLogo.png")}
        style={{
          width: 220,
          resizeMode: "contain"
        }}
      />
      <TouchableOpacity
        onPress={async () => {
          try {
            await NfcManager.registerTagEvent()
          } catch (error) {
            console.warn(error)
            NfcManager.unregisterTagEvent().catch(() => 0)
          }
        }}
      >
        <NeomorphBox
          style={{
            shadowRadius: 3,
            borderRadius: 120,
            backgroundColor: "#fafafa",
            width: 240,
            height: 240,
            justifyContent: "center",
            alignItems: "center"
          }}
        >
          <NeomorphBox
            inner
            style={{
              shadowRadius: 7,
              borderRadius: 100,
              backgroundColor: theme["color-primary-100"],
              width: 200,
              height: 200,
              justifyContent: "center",
              alignItems: "center"
            }}
          >
            <NeomorphBox
              style={{
                shadowRadius: 7,
                borderRadius: 50,
                backgroundColor: "#fafafa",
                width: 100,
                height: 100,
                justifyContent: "center",
                alignItems: "center"
              }}
            >
              <Icon
                name="wifi-outline"
                width={60}
                height={60}
                fill={theme["color-primary-500"]}
              />
            </NeomorphBox>
          </NeomorphBox>
        </NeomorphBox>
      </TouchableOpacity>
      <View
        style={{
          justifyContent: "center",
          alignItems: "center",
          marginBottom: 40,
          padding: 20
        }}
      >
        <Text
          category="h5"
          style={{
            color: theme["color-primary-500"],
            fontWeight: "bold",
            textAlign: "center"
          }}
        >
          Approach your phone near an Aumo device to claim your digital receipt!
        </Text>
      </View>
    </View>
  )
}
