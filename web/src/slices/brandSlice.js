import { createSlice } from '@reduxjs/toolkit'

export const brandSlice = createSlice({
  name: 'brands',
  initialState: {
    value: [],
    loading: false,
    error: null,
  },
  reducers: {
    getData: state => {
      state.loading = true
    },
    getDataSuccess: (state, action) => {
      state.value = action.payload
      state.loading = false
      state.error = null
    },
    getDataFailure: (state) => {
      state.loading = false
      state.error = "Failed to contact network"
    },
  }
})

// actions generated from the slice
export const { getData, getDataSuccess, getDataFailure } = brandSlice.actions

// A selector
export const brandsSelector = (state) => state.value

// The reducer
export default brandSlice.reducer

// Asynchronous thunk action
export function fetchBrands() {
  return async (dispatch) => {
    try {
      const response = await fetch(process.env.REACT_APP_API_URL + 'brands')
      let data = await response.json()
      if (!Array.isArray(data)) {
        data = []
      } 
      dispatch(getDataSuccess(data))
    } catch (error) {
      console.log(error);
      dispatch(getDataFailure())
    }
  }
}
