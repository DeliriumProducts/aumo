import React from "react"
import { ScrollView, StyleSheet, View, Image, Text } from "react-native"

export default function LinksScreen() {
  return (
    <ScrollView style={styles.container}>
      <ScrollView
        style={styles.container}
        contentContainerStyle={styles.contentContainer}
      >
        <View style={styles.welcomeContainer}>
          <Image
            source={
              __DEV__
                ? require("../assets/images/robot-dev.png")
                : require("../assets/images/robot-prod.png")
            }
            style={styles.welcomeImage}
          />
        </View>

        <View style={styles.getStartedContainer}>
          <Text style={styles.getStartedText}>Get started by opening</Text>
          <View
            style={[styles.codeHighlightContainer, styles.homeScreenFilename]}
          ></View>
          <Text style={styles.getStartedText}>
            Change this text and your app will automatically reload.
          </Text>
        </View>
        <View style={styles.helpContainer}></View>
      </ScrollView>
    </ScrollView>
  )
}

LinksScreen.navigationOptions = {
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
