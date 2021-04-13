import React, { useState } from "react"
import { View, TextInput, Button, Image, Text } from "react-native"
import PersonalInfoStyles from "./PersonalInfoStyles";

const PersonalInfo = () => {
    const [name, setName] = useState("");
    const [image, setImage] = useState("");

    return (
        <View style={PersonalInfoStyles.personalInfoContainer}>
            <Image 
                style={PersonalInfoStyles.logo}
                source={require("../assets/favicon.png")}
            />
        
            <View style={PersonalInfoStyles.nameInput}>
                <Text style={PersonalInfoStyles.nameText}>
                    Please enter your name:
                </Text>

                <TextInput style={PersonalInfoStyles.nameTextInput}
                    onChangeText = {(text) => setName(text)}
                    value = {name}
                />
            </View>

            <Button color="red"
                title="Subscribe" 
                onPress={() => {}}
            />

        </View>
    );
};

export default PersonalInfo