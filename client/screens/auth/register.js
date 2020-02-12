import {
  Button,
  Icon,
  Input,
  Layout,
  Modal,
  Spinner,
  Text
} from "@ui-kitten/components"
import aumo from "aumo"
import React, { useState } from "react"
import { Image, ScrollView, StyleSheet, View } from "react-native"
import theme from "../../theme"

export default function RegisterScreen(props) {
  const [email, setEmail] = useState("")
  const [name, setName] = useState("")
  const [password, setPassword] = useState("")
  const [err, setErr] = React.useState("")
  const [showModal, setShowModal] = React.useState(false)
  const [loading, setLoading] = React.useState(false)

  const register = async props => {
    try {
      setLoading(true)
      await aumo.auth.register({
        email: email.trim(),
        name,
        password: password.trim(),
        avatar: "https://i.imgur.com/4Ws6pd9.png"
      })
      setShowModal(true)
    } catch (e) {
      switch (e.response.status) {
        case 400:
          setErr("Bad Request")
          break
        case 500:
          setErr("Internal server error")
          break
      }
    } finally {
      setLoading(false)
    }
  }

  return (
    <ScrollView
      style={styles.container}
      contentContainerStyle={styles.contentContainer}
    >
      <View>
        <View style={styles.mainContainer}>
          <Image
            source={require("../../assets/AumoLogo.png")}
            style={styles.aumo}
          />
          <Text style={styles.subHeading}>The future of receipts.</Text>
        </View>
        <View style={styles.inputform}>
          <Input
            placeholder="Name"
            size="medium"
            icon={style => <Icon {...style} name="person-outline" />}
            value={name}
            onChangeText={setName}
            style={[styles.emailInput, { borderRadius: 10 }]}
          />
          <Input
            placeholder="Email"
            size="medium"
            icon={style => <Icon {...style} name="email-outline" />}
            value={email}
            onChangeText={setEmail}
            style={[styles.emailInput, { borderRadius: 10 }]}
          />
          <Input
            placeholder="Password"
            size="medium"
            secureTextEntry={true}
            icon={style => <Icon {...style} name="lock-outline" />}
            value={password}
            style={{ borderRadius: 10 }}
            onChangeText={setPassword}
          />
          {err != "" && (
            <View style={styles.errorContainer}>
              <Text style={{ color: "white" }}>{err}</Text>
            </View>
          )}
        </View>
      </View>
      <Modal
        visible={showModal}
        backdropStyle={styles.backdrop}
        onBackdropPress={() => {
          setShowModal(false)
          props.navigation.popToTop()
        }}
      >
        <Layout level="3" style={styles.modalContainer}>
          <Text>
            Confirmation email has been sent! Check your email to verify your
            account!
          </Text>
          <Button
            size="small"
            style={{
              marginTop: 10
            }}
            status="success"
            onPress={() => {
              setShowModal(false)
              props.navigation.popToTop()
            }}
          >
            DISMISS
          </Button>
        </Layout>
      </Modal>
      <View
        style={[styles.mainContainer, { paddingRight: 32, paddingLeft: 32 }]}
      >
        <View style={{ marginBottom: 15 }}>
          {loading && <Spinner size="giant" />}
        </View>
        <Button
          disabled={loading}
          style={{ width: "100%", marginBottom: 10, borderRadius: 10 }}
          size="large"
          icon={style => <Icon name="edit-outline" {...style} />}
          onPress={register}
        >
          Register
        </Button>
      </View>
    </ScrollView>
  )
}

const styles = StyleSheet.create({
  container: {
    backgroundColor: theme["color-background-main"],
    flex: 1
  },
  contentContainer: {
    justifyContent: "space-between",
    height: "100%",
    paddingTop: 30
  },
  modalContainer: {
    justifyContent: "center",
    alignItems: "center",
    borderRadius: 8,
    width: 256,
    padding: 16
  },
  backdrop: {
    backgroundColor: "rgba(0, 0, 0, 0.5)"
  },
  mainContainer: {
    alignItems: "center",
    marginTop: 10,
    marginBottom: 20
  },
  aumo: {
    width: 220,
    resizeMode: "contain",
    marginBottom: -20
  },
  getStartedContainer: {
    alignItems: "center",
    marginHorizontal: 50
  },
  subHeading: {
    fontSize: 17,
    color: theme["color-primary-500"],
    marginBottom: 20,
    textAlign: "center"
  },
  inputform: {
    paddingRight: 32,
    paddingLeft: 32
  },
  emailInput: {
    marginBottom: 10
  },
  errorContainer: {
    marginTop: 20,
    borderRadius: 8,
    padding: 15,
    width: "100%",
    backgroundColor: theme["color-danger-500"]
  }
})
