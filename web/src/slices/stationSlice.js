import { createSlice } from '@reduxjs/toolkit'
import GetUrlParams from '../lib/urlParams';

const RADIUS_MIN = 1;
const RADIUS_MAX = 20;

export const stationSlice = createSlice({
  name: 'stations',
  initialState: {
    value: [],
    loading: true,
    error: null, 
    location: [51.4649, -0.1596],
    sortKey: 'distance',
    filters: {
      radius: 3,
      brands: null
    }
  },
  reducers: {
    getData: state => {
      state.loading = true
      state.error = null
    },
    getDataSuccess: (state, action) => {
      state.value = action.payload
      state.loading = false
      state.error = null
    },
    getDataFailure: (state, action) => {
      state.loading = false
      state.error = action.payload
    },
    updateLocation: (state, action) => {
      state.location = action.payload
    },
    updateSort: (state, action) => {
      state.sortKey = action.payload
    },
    updateRadius: (state, action) => {
      let val = parseInt(action.payload)
      if (val < RADIUS_MIN) val = RADIUS_MIN
      if (val > RADIUS_MAX) val = RADIUS_MAX
      state.filters.radius = val;
    },
    updateFilters: (state, action) => {
      for (let [key, value] of Object.entries(action.payload)) {
        state.filters[key] = value
      }
    }
  }
})

// actions generated from the slice
export const { getData, getDataSuccess, getDataFailure, updateLocation, updateSort, updateFilters, updateRadius } = stationSlice.actions

// A selector
export const stationsSelector = (state) => state.value

// The reducer
export default stationSlice.reducer

// Asynchronous thunk action
export function fetchData() {
  return async (dispatch, getState) => {
    try {
      dispatch(getData())
      const pos = await new Promise((resolve, reject) => {
        navigator.geolocation.getCurrentPosition(resolve, reject);
      });

      const location = [pos.coords.latitude, pos.coords.longitude]
      dispatch(updateLocation(location))
    } catch (error) {
      dispatch(getDataFailure("Unable to get location"))
      return
    }

    try {
      const state = getState()
      const response = await fetch(process.env.REACT_APP_API_URL + GetUrlParams({
        latitude: state.stations.location[0],
        longitude: state.stations.location[1],
        radius: state.stations.filters.radius,
        brands: state.stations.filters.brands
      }))

      let data = await response.json()
      if (!Array.isArray(data)) {
        data = []
      } 

      dispatch(getDataSuccess(data))
    } catch (error) {
      dispatch(getDataFailure("Failed to contact network"))
    }
  }
}
