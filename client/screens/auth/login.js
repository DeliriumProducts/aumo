import { Button, Icon, Input, Spinner, Text } from "@ui-kitten/components"
import aumo from "aumo"
import React from "react"
import { Image, ScrollView, StyleSheet, View } from "react-native"
import { TouchableOpacity } from "react-native-gesture-handler"
import { Context } from "../../context/context"
import { actions } from "../../context/providers/provider"
import Routes from "../../navigation/routes"

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
          setErr("Invalid username or password")
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
          <Text style={styles.subheading}>The future of receipts.</Text>
        </View>
        <View style={styles.inputform}>
          <Input
            placeholder="Email"
            icon={style => <Icon {...style} name="email-outline" />}
            value={email}
            onChangeText={setEmail}
            style={[styles.emailInput, { borderRadius: 10 }]}
          />
          <Input
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
            style={{ borderRadius: 10 }}
            onChangeText={setPassword}
          />
          <TouchableOpacity onPress={goToRegister}>
            <Text
              style={[
                styles.subheading,
                {
                  fontSize: 14,
                  textAlign: "right",
                  marginTop: 8,
                  color: "#AAA"
                }
              ]}
            >
              Forgot password?
            </Text>
          </TouchableOpacity>
          {err != "" && (
            <View style={styles.errorContainer}>
              <Text style={{ color: "white" }}>{err}</Text>
            </View>
          )}
        </View>
      </View>
      <View
        style={[styles.mainContainer, { paddingRight: 32, paddingLeft: 32 }]}
      >
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
      </View>
    </ScrollView>
  )
}

LoginScreen.navigationOptions = {
  headerShown: false
}

const styles = StyleSheet.create({
  container: {
    backgroundColor: "#F7F9FC",
    flex: 1
  },
  contentContainer: {
    justifyContent: "space-between",
    height: "100%",
    paddingTop: 30
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
  subheading: {
    fontSize: 17,
    color: "#083AA4",
    // lineHeight: 24,
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
    borderRadius: 4,
    padding: 15,
    width: "100%",
    backgroundColor: "#e9453b"
  }
})
