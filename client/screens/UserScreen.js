import React from "react"
import { ScrollView, StyleSheet, Text, View } from "react-native"

export default function UesrScreen() {
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
  }
})
