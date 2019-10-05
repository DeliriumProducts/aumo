import React, { useState } from "react"
import { Input, Icon, Button } from "react-native-ui-kitten"
import axios from "axios"
import { BACKEND_URL } from "../config"
import * as SecureStore from "expo-secure-store"

import {
  Image,
  Platform,
  ScrollView,
  StyleSheet,
  Text,
  View
} from "react-native"
import { TouchableOpacity } from "react-native-gesture-handler"

export default function LoginScreen(props) {
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")

  const handleLogin = async () => {
    try {
      const res = await axios.post(BACKEND_URL + "/users/login", {
        email,
        password
      })

      if (res.status == 200) {
        const cookie = res.headers["set-cookie"][0]
        await SecureStore.setItemAsync("aumo", cookie)
        props.navigation.navigate("Main")
      }
    } catch (error) {
      console.log(error)
    }
  }

  const goToRegister = () => {
    props.navigation.navigate("Register")
  }

  return (
    <ScrollView
      style={styles.container}
      contentContainerStyle={styles.contentContainer}
    >
      <View>
        <View style={styles.welcomeContainer}>
          <Image
            source={require("../assets/images/AumoLogo.png")}
            style={styles.welcomeImage}
          />
          <Text style={styles.getStartedText}>The future of receipts.</Text>
        </View>
        <View style={styles.inputform}>
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
          <TouchableOpacity onPress={goToRegister}>
            <Text
              style={[
                styles.getStartedText,
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
        </View>
      </View>
      <View
        style={[styles.welcomeContainer, { paddingRight: 32, paddingLeft: 32 }]}
      >
        <Button
          style={{ width: "100%", marginBottom: 10, borderRadius: 10 }}
          size="large"
          state="outline"
          onPress={handleLogin}
        >
          LOGIN
        </Button>
        <TouchableOpacity onPress={goToRegister}>
          <Text style={styles.getStartedText}>Create an account</Text>
        </TouchableOpacity>
      </View>
    </ScrollView>
  )
}

LoginScreen.navigationOptions = {
  header: null
}

const styles = StyleSheet.create({
  container: {
    flex: 1
  },
  developmentModeText: {
    marginBottom: 20,
    color: "rgba(0,0,0,0.4)",
    fontSize: 14,
    lineHeight: 19,
    textAlign: "center"
  },
  contentContainer: {
    justifyContent: "space-between",
    height: "100%",
    paddingTop: 30
  },
  welcomeContainer: {
    alignItems: "center",
    marginTop: 10,
    marginBottom: 20
  },
  welcomeImage: {
    width: 220,
    resizeMode: "contain",
    marginBottom: -20
  },
  getStartedContainer: {
    alignItems: "center",
    marginHorizontal: 50
  },
  homeScreenFilename: {
    marginVertical: 7
  },
  codeHighlightText: {
    color: "rgba(96,100,109, 0.8)"
  },
  codeHighlightContainer: {
    backgroundColor: "rgba(0,0,0,0.05)",
    borderRadius: 3,
    paddingHorizontal: 4
  },
  getStartedText: {
    fontSize: 17,
    color: "#083AA4",
    // lineHeight: 24,
    marginBottom: 20,
    textAlign: "center"
  },
  tabBarInfoContainer: {
    position: "absolute",
    bottom: 0,
    left: 0,
    right: 0,
    ...Platform.select({
      ios: {
        shadowColor: "black",
        shadowOffset: { width: 0, height: -3 },
        shadowOpacity: 0.1,
        shadowRadius: 3
      },
      android: {
        elevation: 20
      }
    }),
    alignItems: "center",
    backgroundColor: "#fbfbfb",
    paddingVertical: 20
  },
  tabBarInfoText: {
    fontSize: 17,
    color: "rgba(96,100,109, 1)",
    textAlign: "center"
  },
  navigationFilename: {
    marginTop: 5
  },
  helpContainer: {
    marginTop: 15,
    alignItems: "center"
  },
  helpLink: {
    paddingVertical: 15
  },
  helpLinkText: {
    fontSize: 14,
    color: "#2e78b7"
  },
  inputform: {
    paddingRight: 32,
    paddingLeft: 32
  },
  emailInput: {
    marginBottom: 10
  }
})
