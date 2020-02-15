import { Button, Icon, Input, Spinner, Text } from "@ui-kitten/components"
import aumo from "aumo"
import React from "react"
import { Image, KeyboardAvoidingView, View } from "react-native"
import { TouchableOpacity } from "react-native-gesture-handler"
import styled from "styled-components/native"
import ErrorContainer from "../../components/ErrorContainer"
import { Context } from "../../context/context"
import { actions } from "../../context/providers/provider"
import Routes from "../../navigation/routes"
import theme from "../../theme"

export default function LoginScreen(props) {
  const [email, setEmail] = React.useState("")
  const [password, setPassword] = React.useState("")
  const [passwordVisible, setPasswordVisible] = React.useState(false)
  const [err, setErr] = React.useState("")
  const [loading, setLoading] = React.useState(false)
  const ctx = React.useContext(Context)

  const handleLogin = async () => {
    try {
      setLoading(true)
      const response = await aumo.auth.login({
        email: email.trim(),
        password: password.trim()
      })

      ctx.dispatch({ type: actions.SET_USER, payload: response })
    } catch (error) {
      switch (error.response.status) {
        case 400:
          setErr("Bad Request")
          break
        case 401:
          setErr("Invalid email or password")
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

  const goToRegister = () => {
    props.navigation.navigate(Routes.Register)
  }

  return (
    <Container enabled behavior="padding">
      <View>
        <MainContainer>
          <Aumo source={require("../../assets/AumoLogo.png")} />
          <Subheading>The future of receipts.</Subheading>
        </MainContainer>
        <Form>
          <FormInput
            placeholder="Email"
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
            value={password}
            onChangeText={setPassword}
          />
          <TouchableOpacity onPress={goToRegister}>
            <Subheading
              style={{
                fontSize: 14,
                textAlign: "right",
                marginTop: 8,
                color: "#AAA"
              }}
            >
              Forgot password?
            </Subheading>
          </TouchableOpacity>
          {err != "" && <ErrorContainer error={err} />}
        </Form>
      </View>
      <MainContainer style={{ paddingRight: 32, paddingLeft: 32 }}>
        <View style={{ marginBottom: 15 }}>
          {loading && <Spinner size="giant" />}
        </View>
        <Button
          disabled={loading}
          icon={style => <Icon name="log-in-outline" {...style} />}
          style={{ width: "100%", marginBottom: 10, borderRadius: 10 }}
          size="large"
          onPress={handleLogin}
        >
          Login
        </Button>
        <Button
          onPress={goToRegister}
          appearance="ghost"
          size="medium"
          disabled={loading}
          style={{ width: "100%", marginBottom: 10, borderRadius: 10 }}
          icon={style => <Icon name="edit-outline" {...style} />}
        >
          Register
        </Button>
      </MainContainer>
    </Container>
  )
}

const Container = styled(KeyboardAvoidingView)`
  background-color: ${theme["color-background-main"]};
  flex: 1;
  height: 100%;
  justify-content: space-between;
`

const MainContainer = styled(View)`
  align-items: center;
  margin-top: 10px;
  margin-bottom: 20px;
`

const Form = styled(View)`
  padding-horizontal: 32px;
`

const Aumo = styled(Image)`
  width: 220px;
  resize-mode: contain;
  margin-bottom: -20px;
`

const Subheading = styled(Text)`
  font-size: 17px;
  color: ${theme["color-primary-500"]};
  margin-bottom: 20px;
  text-align: center;
`

const FormInput = styled(Input)`
  border-radius: 10px;
`
