
import { configureStore } from '@reduxjs/toolkit';
import { Provider } from 'react-redux';

import './App.css';
import Map from './components/Map'
import menuSlice from './slices/menuSlice';
import stationsSlice from './slices/stationSlice'

import { grey } from '@mui/material/colors';
import { CssBaseline, ThemeProvider, createTheme } from '@mui/material';

function App() {
  const theme = createTheme({
    palette: {
      primary: {
        // main: '#D5FC7F'
        main: grey[500]
      },
    },
  });

  const store = configureStore({
    reducer: {
      stations: stationsSlice,
      menus: menuSlice
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
