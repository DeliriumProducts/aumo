import React, { useState } from "react"
import axios from "axios"
import { Input, Icon, Button } from "react-native-ui-kitten"

import { Image, ScrollView, StyleSheet, Text, View } from "react-native"

import { BACKEND_URL } from "../config"

export default function RegisterScreen(props) {
  const [email, setEmail] = useState("")
  const [name, setName] = useState("")
  const [password, setPassword] = useState("")

  const register = async () => {
    try {
      const resp = await axios.post(BACKEND_URL + "/users/register", {
        email,
        name,
        password,
        avatar: "https://i.imgur.com/4Ws6pd9.png"
      })

      if (resp.status === 200) {
        props.navigation.navigate("LogIn")
      }
    } catch (e) {}
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
        </View>
      </View>
      <View
        style={[styles.welcomeContainer, { paddingRight: 32, paddingLeft: 32 }]}
      >
        <Button
          style={{ width: "100%", marginBottom: 10, borderRadius: 10 }}
          size="large"
          state="outline"
          onPress={register}
        >
          REGISTER
        </Button>
      </View>
    </ScrollView>
  )
}

RegisterScreen.navigationOptions = {
  header: null
}

const styles = StyleSheet.create({
  container: {
    flex: 1
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
  getStartedText: {
    fontSize: 17,
    color: "#083AA4",
    marginBottom: 20,
    textAlign: "center"
  },
  inputform: {
    paddingRight: 32,
    paddingLeft: 32
  },
  emailInput: {
    marginBottom: 10
  }
})
