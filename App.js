import { StatusBar } from 'expo-status-bar';
import { StyleSheet, Text, View } from 'react-native';
import DomesticResult from './Pages/DomesticResult';

export default function App() {
  return (
    <View style={styles.container}>
      <DomesticResult></DomesticResult>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});
