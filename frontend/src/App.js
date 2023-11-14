import logo from './logo.svg';
import './App.css';
import Map from './Map'
import { CssBaseline, ThemeProvider, useTheme } from '@mui/material';

function App() {
  const theme = useTheme();

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Map />
  </ThemeProvider>
  );
}

export default App;
