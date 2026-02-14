import { createSelector, createSlice } from '@reduxjs/toolkit'
import { AppDispatch, AppStore, RootState } from '@/store';

import { Station } from '@/models/station';
import { LatLng, LatLngBounds } from 'leaflet';

import GetUrlParams from '@/lib/url_params';
import { sendErrorToast } from '@/components/toast';
import { calculateDistance } from '@/lib/geo_utils';

// const RADIUS_MIN = 1;
// const RADIUS_MAX = 20;

export interface StationStore {
  value: Station[],
  loading: boolean,
  error?: string, 
  location: [number, number],
  bounds?: [number, number, number, number],
  sortKey: string,
  filters: StationFilters,
}

export interface StationFilters {
  brands: string[],
  fuel_types: string[]
}

const initialState: StationStore = {
  value: [],
  loading: true,
  error: undefined, 
  location: [51.4649, -0.1596],
  sortKey: 'distance',
  filters: {
    brands: [],
    fuel_types: [],
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
      // calculate distance using haversine
      action.payload.forEach((station: Station) => {
        const distance = calculateDistance(new LatLng(state.location[0], state.location[1]), new LatLng(station.latitude, station.longitude));
        station.distance = Math.round(distance * 10) / 10; // round to 1 decimal place
      });
      state.value = action.payload
      state.loading = false
      state.error = undefined
    },
    getDataFailure: (state, action) => {
      state.loading = false
      state.error = action.payload
    },
    updateBounds: (state, action) => {
      if (!action.payload) {
        state.bounds = undefined
        return
      }
      state.bounds = action.payload
    },
    updateLocation: (state, action) => {
      state.location = action.payload
    },
    updateSort: (state, action) => {
      state.sortKey = action.payload
    },
    updateFilters: (state, action) => {
      const payload = action.payload as StationFilters;
      state.filters = payload;
    }
  }
})

// actions generated from the slice
export const { getData, getDataSuccess, getDataFailure, updateBounds, updateLocation, updateSort, updateFilters } = stationSlice.actions

// A selector
export const stationsSelector = (state: RootState) => state.stations.value

const bounds = (state: RootState) => state.stations.bounds
export const boundsSelector = createSelector(bounds, (bounds) => bounds ? 
  new LatLngBounds(
    bounds.slice(0, 2) as [number, number], 
    bounds.slice(2, 4) as [number, number]
  ) : undefined)

// The reducer
export default stationSlice.reducer

// thunks
export function fetchData(bounds: LatLngBounds) {
  return async (dispatch: AppDispatch, getState: AppStore['getState']) => {
    try {
      dispatch(getData())
      dispatch(updateBounds([
        bounds.getNorthWest().lat, bounds.getNorthWest().lng, 
        bounds.getSouthEast().lat, bounds.getSouthEast().lng
      ]))

      const state = getState()
      const response = await fetch(import.meta.env.VITE_APP_API_URL + 'stations/query' + GetUrlParams({
        coords: [bounds.getNorthWest().lng, bounds.getNorthWest().lat, bounds.getSouthEast().lng, bounds.getSouthEast().lat],
        ...state.stations.filters,
      }))

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || "Failed to get stations");
      }

      let data = await response.json()
      if (!Array.isArray(data)) {
        data = []
      } 

      dispatch(getDataSuccess(data))
    } catch (error) {
      console.error(error)
      dispatch(getDataFailure("Failed to get stations"))
      sendErrorToast((error as Error).message, "Failed to get stations")
    }
  }
}

export function loadCurrentPosition() {
  return async (dispatch: AppDispatch) => {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition((position) => {
        dispatch(updateLocation([position.coords.latitude, position.coords.longitude]))
      });
    }
  }
}