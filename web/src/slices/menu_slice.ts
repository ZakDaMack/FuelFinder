import { RootState } from '@/store';
import { createSlice } from '@reduxjs/toolkit'

interface MenuStore {
  stations: boolean;
  preferences: boolean;
  tooltips: boolean;
  info: boolean;
}

const initialState: MenuStore = {
  stations: false,
  preferences: false,
  tooltips: false,
  info: false,
}

export const menuSlice = createSlice({
  name: 'menus',
  initialState,
  reducers: {
    openMenu: (state, action) => {
      state[action.payload as keyof MenuStore] = true
    },
    closeMenu: (state, action) => {
      state[action.payload as keyof MenuStore] = false
    },
    closeAll: (state) => {
      Object.keys(state).forEach(key => 
        state[key as keyof MenuStore] = false
      )
    },
  }
})

// actions generated from the slice
export const { openMenu, closeMenu, closeAll } = menuSlice.actions

// A selector
export const brandsSelector = (state: RootState) => state.brands.value

// The reducer
export default menuSlice.reducer