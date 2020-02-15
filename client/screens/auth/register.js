import {
  Button,
  Icon,
  Layout,
  Modal,
  Spinner,
  Text
} from "@ui-kitten/components"
import aumo from "aumo"
import React, { useState } from "react"
import { StyleSheet, View } from "react-native"
import styled from "styled-components/native"
import ErrorContainer from "../../components/ErrorContainer"
import {
  Aumo,
  Container,
  Form,
  FormInput,
  MainContainer,
  Subheading
} from "./components"

export default function RegisterScreen(props) {
  const [email, setEmail] = useState("")
  const [name, setName] = useState("")
  const [password, setPassword] = useState("")
  const [passwordVisible, setPasswordVisible] = React.useState(false)
  const [err, setErr] = React.useState("")
  const [showModal, setShowModal] = React.useState(false)
  const [loading, setLoading] = React.useState(false)

  const register = async () => {
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
        case 422:
          setErr("Email already exists")
          break
        case 500:
          setErr("Internal server error")
          break
      }
    } finally {
      setLoading(false)
    }
  }

  const onPasswordIconPress = () => {
    setPasswordVisible(!passwordVisible)
  }

  return (
    <Container>
      <View>
        <MainContainer>
          <Aumo source={require("../../assets/AumoLogo.png")} />
          <Subheading>The future of receipts.</Subheading>
        </MainContainer>
        <Form>
          <FormInput
            placeholder="Name"
            size="medium"
            icon={style => <Icon {...style} name="person-outline" />}
            value={name}
            onChangeText={setName}
            style={{ marginBottom: 10 }}
          />
          <FormInput
            placeholder="Email"
            size="medium"
            icon={style => <Icon {...style} name="email-outline" />}
            value={email}
            onChangeText={setEmail}
            style={{ marginBottom: 10 }}
          />
          <FormInput
            placeholder="Password"
            secureTextEntry={!passwordVisible}
            icon={style => (
              <Icon
                {...style}
                name={passwordVisible ? "eye-outline" : "eye-off-outline"}
              />
            )}
            onIconPress={onPasswordIconPress}
            style={{ marginBottom: 10 }}
            value={password}
            onChangeText={setPassword}
          />
          {err != "" && <ErrorContainer error={err} />}
        </Form>
      </View>
      <Modal
        visible={showModal}
        backdropStyle={styles.backdrop}
        onBackdropPress={() => {
          setShowModal(false)
          props.navigation.popToTop()
        }}
      >
        <ModalContainer level="3">
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
        </ModalContainer>
      </Modal>
      <MainContainer style={{ paddingRight: 32, paddingLeft: 32 }}>
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
      </MainContainer>
    </Container>
  )
}

const ModalContainer = styled(Layout)`
  justify-content: center;
  align-items: center;
  border-radius: 8px;
  width: 256px;
  padding: 16px;
`

const styles = StyleSheet.create({
  backdrop: {
    backgroundColor: "rgba(0, 0, 0, 0.5)"
  }
})
