import axios from "axios"
import * as SecureStore from "expo-secure-store"
import React from "react"
import {
  ScrollView,
  StyleSheet,
  Text,
  View,
  ActivityIndicator
} from "react-native"
import { BACKEND_URL } from "../config"
import { Avatar } from "react-native-ui-kitten"
import { withNavigationFocus } from "react-navigation"

const UserScreen = props => {
  const [user, setUser] = React.useState(null)
  const [loading, setLoading] = React.useState(true)

  React.useEffect(() => {
    if (props.isFocused) {
      fetchUser()
    }
  }, [props.isFocused])

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
    setLoading(false)
  }

  return (
    <View style={styles.container}>
      <ScrollView
        style={styles.container}
        contentContainerStyle={styles.contentContainer}
      >
        {loading ? (
          <ActivityIndicator />
        ) : (
          <View>
            <Text>{user.name}</Text>
            <Avatar source={{ uri: user.avatar }} />
          </View>
        )}
      </ScrollView>
    </View>
  )
}

UserScreen.navigationOptions = {
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

export default withNavigationFocus(UserScreen)
