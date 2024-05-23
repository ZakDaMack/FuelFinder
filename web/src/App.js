
import { configureStore } from '@reduxjs/toolkit';
import { Provider } from 'react-redux';

import './App.css';
import Map from './pages/Map'
import IntroDialog from './pages/IntroDialog';
import LoadingOverlay from './pages/LoadingOverlay';
import Init from './components/Init'
import ErrorSnackbar from './components/ErrorSnackbar';

import menuSlice from './slices/menuSlice';
import stationsSlice from './slices/stationSlice'
import brandSlice from './slices/brandSlice';

import { CssBaseline, ThemeProvider, createTheme } from '@mui/material';

function App() {
  const theme = createTheme({
    palette: {
      primary: {
        main: '#66C745',
        dark: '#00568A'
        // main: grey[500]
      },
      secondary: {
        main: '#ADA9A4'
      },
      black: {
        main: '#000000',
        light: '#E9DB5D',
        dark: '#A29415',
        contrastText: '#FFFFFF',
      },
    },
    typography: {
      fontFamily: '"inter", sans-serif'
    }
  });

  const store = configureStore({
    reducer: {
      stations: stationsSlice,
      menus: menuSlice,
      brands: brandSlice,
    }
  })

  return (
    <Provider store={store}>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <IntroDialog />
        <LoadingOverlay />
        <Init />
        <Map />
        <ErrorSnackbar />
      </ThemeProvider>
    </Provider>
  );
}

export default App;
