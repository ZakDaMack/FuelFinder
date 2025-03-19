import { createSlice } from '@reduxjs/toolkit'

import { Station } from '@/models/station';
import GetUrlParams from '@/lib/url_params';
import { AppDispatch, AppStore, RootState } from '@/store';

const RADIUS_MIN = 1;
const RADIUS_MAX = 20;

export interface StationStore {
  value: Station[],
  loading: boolean,
  error?: string, 
  location: [number, number],
  sortKey: string,
  filters: StationFilters,
}

export interface StationFilters {
  radius: number,
  brands?: string[],
  fuel_types?: string[]
}

const initialState: StationStore = {
  value: [],
  loading: true,
  error: undefined, 
  location: [51.4649, -0.1596],
  sortKey: 'distance',
  filters: {
    radius: 3,
    brands: undefined,
    fuel_types: undefined,
  }
};

export const stationSlice = createSlice({
  name: 'stations',
  initialState,
  reducers: {
    getData: state => {
      state.loading = true
      state.error = undefined
    },
    getDataSuccess: (state, action) => {
      state.value = action.payload
      state.loading = false
      state.error = undefined
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
      const payload = action.payload as StationFilters;
      state.filters = payload;
    }
  }
})

// actions generated from the slice
export const { getData, getDataSuccess, getDataFailure, updateLocation, updateSort, updateFilters, updateRadius } = stationSlice.actions

// A selector
export const stationsSelector = (state: RootState) => state.stations.value

// The reducer
export default stationSlice.reducer

// Asynchronous thunk action
export function fetchData() {
  return async (dispatch: AppDispatch, getState: AppStore['getState']) => {
    try {
      dispatch(getData())
      const pos: GeolocationPosition = await new Promise((resolve, reject) => {
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
      const response = await fetch(import.meta.env.VITE_APP_API_URL + GetUrlParams({
        latitude: state.stations.location[0],
        longitude: state.stations.location[1],
        ...state.stations.filters,
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
