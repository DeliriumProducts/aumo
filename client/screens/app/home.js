import { Icon, Text, Button, Spinner } from "@ui-kitten/components"
import aumo from "aumo"
import React from "react"
import { Image, View } from "react-native"
import { TouchableOpacity } from "react-native-gesture-handler"
import { NeomorphBox } from "react-native-neomorph-shadows"
import NfcManager, { NdefParser, NfcEvents } from "react-native-nfc-manager"
import Modal from "../../components/Modal"
import theme from "../../theme"

export default () => {
  const [loading, setLoading] = React.useState(false)
  const [showModal, setShowModal] = React.useState(false)

  React.useEffect(() => {
    NfcManager.start()
    NfcManager.setEventListener(NfcEvents.DiscoverTag, async tag => {
      let msgs = tag.ndefMessage.map(NdefParser.parseText)
      if (msgs.length >= 1) {
        let receiptId = msgs[0]
        setLoading(true)
        try {
          await aumo.receipt.claimReceipt(receiptId)
        } catch (error) {
        } finally {
          setLoading(false)
          setShowModal(true)
        }
      }
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
      <Modal visible={showModal} onBackdropPress={() => setShowModal(true)}>
        <View style={{ width: 256 }}>
          <Text>
            You successfully claimed a receipt and were awarded 500 points!
          </Text>
          <Button
            size="small"
            style={{
              marginTop: 10
            }}
            status="success"
            onPress={() => setShowModal(false)}
          >
            DISMISS
          </Button>
        </View>
      </Modal>
      <Modal visible={loading} onBackdropPress={() => {}}>
        {loading && <Spinner size="giant" />}
      </Modal>
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
