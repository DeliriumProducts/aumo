import React, { useState } from "react"
import { Input, Icon, Button } from "react-native-ui-kitten"

import {
  Image,
  Platform,
  ScrollView,
  StyleSheet,
  Text,
  View
} from "react-native"
import { TouchableOpacity } from "react-native-gesture-handler"

export default function RegisterScreen() {
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")

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
          <TouchableOpacity>
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
        >
          LOGIN
        </Button>
        <TouchableOpacity>
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
