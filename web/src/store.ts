import { configureStore } from "@reduxjs/toolkit";
import { TypedUseSelectorHook, useDispatch, useSelector } from "react-redux";

import menuSlice from './slices/menu_slice'
import brandSlice from './slices/brand_slice'
import detailsSlice from './slices/details_slice'
import stationsSlice from './slices/station_slice'

const store = configureStore({
    reducer: {
      stations: stationsSlice,
      menus: menuSlice,
      brands: brandSlice,
      details: detailsSlice,
    }
})

export { store };  

// Get the type of our store variable
export type AppStore = typeof store

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<AppStore['getState']>

// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = AppStore['dispatch']

// Use throughout app
export const useAppDispatch: () => AppDispatch = useDispatch
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector