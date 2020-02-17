import { Button, Icon, Spinner } from "@ui-kitten/components"
import aumo from "aumo"
import React from "react"
import { useForm } from "react-hook-form"
import { View } from "react-native"
import { TouchableOpacity } from "react-native-gesture-handler"
import ErrorContainer from "../../components/ErrorContainer"
import FormInput from "../../components/FormInput"
import { Context } from "../../context/context"
import { actions } from "../../context/providers/provider"
import Routes from "../../navigation/routes"
import { Aumo, Container, Form, MainContainer, Subheading } from "./components"

export default function LoginScreen(props) {
  const { register, handleSubmit, errors, setValue } = useForm()
  const [passwordVisible, setPasswordVisible] = React.useState(false)
  const [err, setErr] = React.useState("")
  const [loading, setLoading] = React.useState(false)
  const ctx = React.useContext(Context)

  const handleLogin = async data => {
    try {
      setLoading(true)
      const response = await aumo.auth.login({
        email: data.email.trim(),
        password: data.password.trim()
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
        case 404:
          setErr("User doesn't exist")
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
    <Container>
      <View>
        <MainContainer>
          <Aumo source={require("../../assets/AumoLogo.png")} />
          <Subheading>The future of receipts.</Subheading>
        </MainContainer>
        <Form>
          <FormInput
            status={errors.email ? "danger" : "basic"}
            placeholder="Email"
            icon={style => <Icon {...style} name="email-outline" />}
            disabled={loading}
            ref={register("email", {
              required: "Required",
              pattern: {
                value: /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i,
                message: "Must be an email"
              }
            })}
            caption={errors.email?.message}
            onChangeText={val => setValue("email", val)}
          />
          <FormInput
            status={errors.password ? "danger" : "basic"}
            placeholder="Password"
            secureTextEntry={!passwordVisible}
            disabled={loading}
            onIconPress={onPasswordIconPress}
            style={{ marginTop: 8 }}
            icon={style => (
              <Icon
                {...style}
                name={passwordVisible ? "eye-outline" : "eye-off-outline"}
              />
            )}
            ref={register("password", { required: "Required" })}
            onChangeText={val => setValue("password", val)}
            caption={errors.password?.message}
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
          onPress={handleSubmit(handleLogin)}
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
