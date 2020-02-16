import {
  Button,
  Icon,
  Layout,
  Modal,
  Spinner,
  Text
} from "@ui-kitten/components"
import aumo from "aumo"
import React from "react"
import { useForm } from "react-hook-form"
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
  const { register, handleSubmit, errors, setValue } = useForm()
  const [passwordVisible, setPasswordVisible] = React.useState(false)
  const [err, setErr] = React.useState("")
  const [showModal, setShowModal] = React.useState(false)
  const [loading, setLoading] = React.useState(false)

  const handleRegister = async data => {
    try {
      setLoading(true)
      await aumo.auth.register({
        email: data.email.trim(),
        name: data.name.trim(),
        password: data.password.trim(),
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
            status={errors.name ? "danger" : "basic"}
            placeholder="Name"
            size="medium"
            icon={style => <Icon {...style} name="person-outline" />}
            ref={register("name", { required: "Required" })}
            onChangeText={val => setValue("name", val)}
            style={{ marginBottom: 10 }}
          />
          {errors.name && (
            <ErrorContainer
              error={errors.name.message}
              style={{ marginBottom: 10 }}
            />
          )}
          <FormInput
            status={errors.email ? "danger" : "basic"}
            placeholder="Email"
            size="medium"
            icon={style => <Icon {...style} name="email-outline" />}
            ref={register("email", {
              required: "Required",
              pattern: {
                value: /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i,
                message: "Must be an email"
              }
            })}
            style={{ marginBottom: 10 }}
            onChangeText={val => setValue("email", val)}
          />
          {errors.email && (
            <ErrorContainer
              error={errors.email.message}
              style={{ marginBottom: 10 }}
            />
          )}
          <FormInput
            status={errors.password ? "danger" : "basic"}
            placeholder="Password"
            secureTextEntry={!passwordVisible}
            icon={style => (
              <Icon
                {...style}
                name={passwordVisible ? "eye-outline" : "eye-off-outline"}
              />
            )}
            ref={register("password", {
              required: "Required",
              maxLength: {
                value: 24,
                message: "Must be shorter than 24 chars"
              },
              minLength: {
                value: 6,
                message: "Must be longer than 6 chars"
              }
            })}
            onIconPress={onPasswordIconPress}
            style={{ marginBottom: 10 }}
            onChangeText={val => setValue("password", val)}
          />
          {errors.password && (
            <ErrorContainer error={errors.password.message} />
          )}
          {err != "" && <ErrorContainer error={err} />}
        </Form>
      </View>
      <Modal
        visible={showModal}
        backdropStyle={styles.backdrop}
        onBackdropPress={() => {
          props.navigation.popToTop()
          setShowModal(false)
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
          onPress={handleSubmit(handleRegister)}
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
