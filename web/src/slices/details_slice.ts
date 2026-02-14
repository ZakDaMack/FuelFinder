import { createSlice } from '@reduxjs/toolkit'
import { AppDispatch, AppStore, RootState } from '@/store'

import { Station } from '@/models/station';

interface DetailsStore {
  stations: { [key: string]: Station[] };
  loading: boolean;
  error?: string;
}

const initialState: DetailsStore = {
  stations: {},
  loading: false,
  error: undefined,
}

export const detailsStore = createSlice({
  name: 'details',
  initialState,
  reducers: {
    getData: state => { 
      state.loading = true
    },
    getDataSuccess: (state, action) => {
      state.stations[action.payload.id] = action.payload.data
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
export const { getData, getDataSuccess, getDataFailure } = detailsStore.actions

// A selector
export const detailsSelector = (state: RootState, id: string) => state.details.stations[id] || []
export const detailsLoadingSelector = (state: RootState) => state.details.loading
export const detailsErrorSelector = (state: RootState) => state.details.error

// The reducer
export default detailsStore.reducer

// Asynchronous thunk action
export function fetchStations(id: string) {
  return async (dispatch: AppDispatch, getState: AppStore['getState']) => {
    const state = getState();
    const data = state.details.stations[id];
    if (data) return; // already have data, dont fetch it again

    // dont have data, fetch it
    try {
      dispatch(getData()) // set loading state
      const response = await fetch(import.meta.env.VITE_APP_API_URL + `stations/${id}`)
      let data = await response.json()

      if (!response.ok) throw new Error(data.message || "Network response was not ok");
      dispatch(getDataSuccess({ id, data }))
    } catch (error) {
      console.log(error);
      dispatch(getDataFailure())
    }
  }
}
