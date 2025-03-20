import { StrictMode } from 'react'
import { Provider } from 'react-redux'
import { createRoot } from 'react-dom/client'
import { Toaster } from "@/components/ui/sonner"

import { store } from './store'
import Map from './components/map'
import Init from './components/init'
import Preferences from './components/preferences'
import StationList from './components/station_list'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <Provider store={store}>
      <Init />
      <Map />
      <Preferences />
      <StationList />
    </Provider>
    <Toaster />
  </StrictMode>,
)
