import React from "react"
import { ActivityIndicator, StatusBar, View } from "react-native"
import { AuthAPI } from "../api"

import * as SecureStore from "expo-secure-store"

export default class AuthLoadingScreen extends React.Component {
  componentDidMount() {
    this._bootstrapAsync()
  }

  _bootstrapAsync = async () => {
    try {
      await AuthAPI.me()
      this.props.navigation.navigate("Main")
    } catch (e) {
      if (e.response.status === 401) {
        this.props.navigation.navigate("Auth")
      }
    }
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
