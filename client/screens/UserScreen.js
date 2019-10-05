import axios from "axios"
import * as SecureStore from "expo-secure-store"
import React from "react"
import { ScrollView, StyleSheet, Text, View } from "react-native"
import { BACKEND_URL } from "../config"

export default function UesrScreen() {
  const [user, setUser] = React.useState(null)
  React.useEffect(() => {
    fetchUser()
  }, [])

  const fetchUser = async () => {
    try {
      const u = await SecureStore.getItemAsync("aumo")
      const res = await axios.get(BACKEND_URL + "/me", {
        headers: {
          Cookie: u
        },
        withCredentials: true
      })
      setUser(res.data)
    } catch (e) {}
  }

  return (
    <View style={styles.container}>
      <ScrollView
        style={styles.container}
        contentContainerStyle={styles.contentContainer}
      >
        <Text>Epiiic</Text>
      </ScrollView>
    </View>
  )
}

UesrScreen.navigationOptions = {
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
  }
})
