import React from "react"
import {BACKEND_URL} from "../config"
import axios from "axios"
import {
    Image,
    Platform,
    TouchableWithoutFeedback,
    ScrollView,
    StyleSheet,
    Text,
    TouchableOpacity,
    View
} from "react-native"

import {MonoText} from "../components/StyledText"
import Circle from "../components/Circle"
import {AntDesign} from "@expo/vector-icons"

export default function HomeScreen() {
    return (
        <View style={styles.container}>
            <ScrollView
                style={styles.container}
                contentContainerStyle={styles.contentContainer}
            >
                <View style={styles.welcomeContainer}>
                    <Image
                        source={require("../assets/images/AumoLogo.png")}
                        style={styles.welcomeImage}
                    />
                    <Text style={styles.getStartedText}>
                        Approach phone to an aumo device, to get your digital receipt
          </Text>
                </View>
                <TouchableWithoutFeedback
                    onPress={async () => {
                        // const receipt = await axios.post(BACKEND_URL + "/receipts", {
                        //   content: ""
                        // })
                        // const resp = await axios.post(
                        //   BACKEND_URL + "/users/claim-receipt/" + receipt.data.id
                        // )
                    }}
                >
                    <Circle size={240}>
                        <Circle size={200}>
                            <View style={{alignItems: "center"}}>
                                <AntDesign name="wifi" size={60} color="#083AA4" />
                            </View>
                        </Circle>
                    </Circle>
                </TouchableWithoutFeedback>
                <View style={styles.helpContainer}></View>
            </ScrollView>
        </View>
    )
}

HomeScreen.navigationOptions = {
    header: null
}

const styles = StyleSheet.create({
    container: {
        backgroundColor: "#F7F9FC",
        flex: 1
    },
    contentContainer: {
        justifyContent: "space-between",
        alignItems: "center",
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
        paddingHorizontal: 18,
        fontSize: 18,
        fontWeight: "700",
        color: "#083AA4",
        textTransform: "uppercase",
        marginTop: 50,
        textAlign: "center"
    }
})
