import React from "react"
import { ActivityIndicator, StatusBar, View } from "react-native"

import * as SecureStore from "expo-secure-store"

export default class AuthLoadingScreen extends React.Component {
  componentDidMount() {
    this._bootstrapAsync()
  }

  _bootstrapAsync = async () => {
    // TODO: repalce with request to /me and check for status code
    const userSession = await SecureStore.getItemAsync("aumo")

    this.props.navigation.navigate(userSession ? "Main" : "Auth")
  }

  render() {
    return (
      <View>
        <ActivityIndicator />
        <StatusBar barStyle="default" />
      </View>
    )
  }
}
