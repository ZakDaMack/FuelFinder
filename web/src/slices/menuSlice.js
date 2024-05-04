import { createSlice } from '@reduxjs/toolkit'

export const menuSlice = createSlice({
  name: 'menus',
  initialState: {
    stations: false,
    preferences: false,
    tooltips: false,
  },
  reducers: {
    openMenu: (state, action) => {
      state[action.payload] = true
    },
    closeMenu: (state, action) => {
      state[action.payload] = false
    },
    closeAll: (state) => {
      Object.keys(state).forEach(key => 
        state[key] = false
      )
    },
  }
})

// actions generated from the slice
export const { openMenu, closeMenu, closeAll } = menuSlice.actions

// The reducer
export default menuSlice.reducer