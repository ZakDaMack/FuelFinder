import { StrictMode } from 'react'
import { Provider } from 'react-redux'
import { createRoot } from 'react-dom/client'
import { Toaster } from "@/components/ui/sonner"

import { store } from './store'
import Map from './components/map'
import Init from './components/init'
import Preferences from './components/preferences'
import StationList from './components/station_list'
import IntroDialog from './components/intro_dialog'

import LoginPage from './pages/auth/login_form'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <Provider store={store}>
      <IntroDialog />
      <Preferences />
      <StationList />
      <Init />
      <Map />
    </Provider>
    {/* <LoginPage /> */}
    <Toaster />
  </StrictMode>,
)
