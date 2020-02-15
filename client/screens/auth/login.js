import { Button, Icon, Spinner } from "@ui-kitten/components"
import aumo from "aumo"
import React from "react"
import { useForm } from "react-hook-form"
import { View } from "react-native"
import { TouchableOpacity } from "react-native-gesture-handler"
import ErrorContainer from "../../components/ErrorContainer"
import { Context } from "../../context/context"
import { actions } from "../../context/providers/provider"
import Routes from "../../navigation/routes"
import {
  Aumo,
  Container,
  Form,
  FormInput,
  MainContainer,
  Subheading
} from "./components"

export default function LoginScreen(props) {
  const { register, handleSubmit, errors, setValue } = useForm()
  const [passwordVisible, setPasswordVisible] = React.useState(false)
  const [err, setErr] = React.useState("")
  const [loading, setLoading] = React.useState(false)
  const ctx = React.useContext(Context)

  const handleLogin = async data => {
    console.log("uh ok")
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
            placeholder="Email"
            icon={style => <Icon {...style} name="email-outline" />}
            ref={register("email", {
              required: "Email is required",
              pattern: {
                value: /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i,
                message: "Must be an email"
              }
            })}
            onChangeText={val => setValue("email", val)}
            style={{ marginBottom: 10 }}
          />
          {errors.email && (
            <ErrorContainer
              error={errors.email.message}
              style={{ marginBottom: 10 }}
            />
          )}
          <FormInput
            placeholder="Password"
            secureTextEntry={!passwordVisible}
            onIconPress={onPasswordIconPress}
            icon={style => (
              <Icon
                {...style}
                name={passwordVisible ? "eye-outline" : "eye-off-outline"}
              />
            )}
            ref={register("password", { required: "Password is required" })}
            onChangeText={val => setValue("password", val)}
          />
          {errors.password && (
            <ErrorContainer
              error={errors.password.message}
              style={{ marginTop: 10 }}
            />
          )}
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
