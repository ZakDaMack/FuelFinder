
import { configureStore } from '@reduxjs/toolkit';
import { Provider } from 'react-redux';
import './App.css';
import Map from './Map'
import { CssBaseline, ThemeProvider, createTheme } from '@mui/material';
import stationsSlice from './slices/stationSlice'

function App() {
  const theme = createTheme({
    palette: {
      primary: {
        main: '#D5FC7F'
      },
    },
  });

  const store = configureStore({
    reducer: {
      stations: stationsSlice
    }
  })

  return (
    <Provider store={store}>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <Map />
      </ThemeProvider>
    </Provider>
  );
}

export default App;
