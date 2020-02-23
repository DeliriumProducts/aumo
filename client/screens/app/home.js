import { Icon, Text, Button, Spinner } from "@ui-kitten/components"
import aumo from "aumo"
import React from "react"
import { Image, View, Alert } from "react-native"
import { TouchableOpacity } from "react-native-gesture-handler"
import { NeomorphBox } from "react-native-neomorph-shadows"
import NfcManager, {
  NdefParser,
  NfcEvents,
  Ndef,
  NfcTech
} from "react-native-nfc-manager"
import ndef from "react-native-nfc-manager/ndef-lib/index"
import Modal from "../../components/Modal"
import theme from "../../theme"

export default () => {
  const [loading, setLoading] = React.useState(false)
  const [showModal, setShowModal] = React.useState(false)
  const [listening, setListening] = React.useState(false)
  const [text, setText] = React.useState("")

  React.useEffect(() => {
    NfcManager.start()
    NfcManager.setEventListener(NfcEvents.DiscoverTag, async tag => {
      let msgs = tag.ndefMessage.map(NdefParser.parseText)
      console.warn(msgs)
      if (msgs.length > 1) {
        let rid = msgs.shift()
        setLoading(true)
        try {
          await NfcManager.requestTechnology(NfcTech.Ndef)
          NfcManager.writeNdefMessage(
            Ndef.encodeMessage(msgs.map(m => Ndef.textRecord(m)))
          )
          await NfcManager.cancelTechnologyRequest()
        } catch (e) {
          setText(e.toString())
          await NfcManager.cancelTechnologyRequest()
        }
        try {
          await aumo.receipt.claimReceipt(rid)
          setShowModal(true)
          setText(
            "You successfully claimed a receipt and were rewarded 500 points!"
          )
        } catch (e) {
          setShowModal(true)
          setText(e.response.data.error)
        } finally {
          setLoading(false)
        }
      } else if (msgs[0] == ".") {
        setShowModal(true)
        setText("There is no receipt to be claimed.")
      }
      NfcManager.unregisterTagEvent().catch(() => 0)
      setListening(false)
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
            setListening(true)
          } catch (ex) {
            console.warn("ex", ex)
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
          <Text>{text}</Text>
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
        <Text
          style={{
            color: theme["color-primary-500"],
            fontWeight: "bold",
            textAlign: "center"
          }}
        >
          {listening && "Listening ..."}
        </Text>
      </View>
    </View>
  )
}
