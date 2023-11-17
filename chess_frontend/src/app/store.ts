import { configureStore, ThunkAction, Action } from "@reduxjs/toolkit";

import globalReducer from "../features/public";

export type TReducers = {
  global: typeof globalReducer;
};

export const store = configureStore({
  reducer: {
    global: globalReducer,
  } as TReducers,
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
