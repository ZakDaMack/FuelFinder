import { AppDispatch, RootState } from '@/store'
import { createSlice } from '@reduxjs/toolkit'

interface BrandStore {
  value: string[];
  loading: boolean;
  error?: string;
}

const initialState: BrandStore = {
  value: [],
  loading: false,
  error: undefined,
}

export const brandSlice = createSlice({
  name: 'brands',
  initialState,
  reducers: {
    getData: state => {
      state.loading = true
    },
    getDataSuccess: (state, action) => {
      state.value = action.payload
      state.loading = false
      state.error = undefined
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
export const brandsSelector = (state: RootState) => state.brands.value

// The reducer
export default brandSlice.reducer

// Asynchronous thunk action
export function fetchBrands() {
  return async (dispatch: AppDispatch) => {
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
