import { createSlice } from '@reduxjs/toolkit'

export const stationSlice = createSlice({
  name: 'stations',
  initialState: {
    value: [],
    loading: false,
    error: null, 
    location: [51.4649, -0.1596],
    filters: {
      radius: 3
    }
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
    updateLocation: (state, action) => {
      state.location = action.payload
    },
    updateSettings: (state, action) => {
      for (let [key, value] of Object.entries(action.payload)) {
        state.filters[key] = value
      }
    }
  }
})

// actions generated from the slice
export const { getData, getDataSuccess, getDataFailure, updateLocation, updateSettings } = stationSlice.actions

// A selector
export const stationsSelector = (state) => state.value

// The reducer
export default stationSlice.reducer

// Asynchronous thunk action
export function fetchData(location) {
  return async (dispatch, getState) => {
    dispatch(getData())
    dispatch(updateLocation(location))
    const state = getState()
    
    try {
      const response = await fetch(process.env.REACT_APP_API_URL + '?' + new URLSearchParams({
        latitude: state.stations.location[0],
        longitude: state.stations.location[1],
        radius: state.stations.filters.radius
      }))

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
