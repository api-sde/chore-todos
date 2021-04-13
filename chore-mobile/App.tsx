import { StatusBar } from 'expo-status-bar';
import React from 'react';
import { SafeAreaView, StyleSheet, Text, View } from 'react-native';
import PersonalInfo from './components/PersonalInfo';
import AppStyles from './GenericStyles/AppStyles';

export default function App() {
  return (
    <SafeAreaView style={AppStyles.container}>
     <PersonalInfo/>
     <StatusBar style="auto" />
    </SafeAreaView>
  );
};
