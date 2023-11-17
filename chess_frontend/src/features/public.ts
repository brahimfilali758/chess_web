import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { RootState } from "../app/store";

export interface GlobalState {
  themeMode: "light" | "dark";
}

const initialState: GlobalState = {
  themeMode: "dark",
};

export const globalSlice = createSlice({
  name: "global",
  initialState: initialState,
  reducers: {
    setThemeMode: (
      state: GlobalState,
      action: PayloadAction<"light" | "dark">
    ) => {
      state.themeMode = action.payload;
    },
  },
});

export const { setThemeMode } = globalSlice.actions;

export const selectThemeMode = (state: RootState) => state.global.themeMode;
export default globalSlice.reducer;
