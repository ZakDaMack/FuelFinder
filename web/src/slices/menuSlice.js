import { createSlice } from '@reduxjs/toolkit'

export const menuSlice = createSlice({
  name: 'menus',
  initialState: {
    stations: false,
    preferences: false,
    tooltips: false,
  },
  reducers: {
    // get menu you want opened and set to true, close all toher menus
    updateMenu: (state, action) => {
        Object.keys(state).forEach(key => 
            state[key] = (action.payload === key)
        )
    },
    closeAll: (state) => {
        Object.keys(state).forEach(key => 
            state[key] = false
        )
    },
  }
})

// actions generated from the slice
export const { updateMenu, closeAll } = menuSlice.actions

// The reducer
export default menuSlice.reducer