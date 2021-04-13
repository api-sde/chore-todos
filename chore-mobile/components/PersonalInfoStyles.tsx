import { StyleSheet } from "react-native"

export default StyleSheet.create({
    container: {
      flex: 1,
      backgroundColor: '#007AFF',
    },

    personalInfoContainer: {
        flex: 1,
        padding: 10,
        justifyContent: "space-between",
        alignItems: "stretch",
    },

    logo: { width: "auto", resizeMode: "contain" },

    nameInput: {
        alignSelf: "center",
    },

    nameText: {
        fontSize: 20,
    },

    nameTextInput: {
        borderColor: "rgba(52, 52, 52, 0.8)",
        borderWidth: 1,
        borderRadius: 4,
        fontSize: 18,
    },

    submitButton: {
    }

  });