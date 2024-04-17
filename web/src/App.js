import './App.css';
import Map from './Map'
import { CssBaseline, ThemeProvider, createTheme } from '@mui/material';

function App() {
  const theme = createTheme({
    palette: {
      primary: {
        main: '#D5FC7F'
      },
    },
  });

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Map />
    </ThemeProvider>
  );
}

export default App;
