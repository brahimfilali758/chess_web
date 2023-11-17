import React, { useMemo, useState } from "react";
import logo from "./logo.svg";
import styles from "./App.module.css";
import Button from "@mui/material/Button";
import ApplicationToolbar from "./Components/ApplicationToolBar";
import { ThemeProvider } from "@emotion/react";
import { createTheme } from "@mui/material";
import * as lightMode from "./themeLight";
import * as darkMode from "./themeDark";
import { Chessboard } from "react-chessboard";

function App() {
  const [themeMode, setThemeMode] = useState<"light" | "dark">("light");

  const theme = useMemo(() => {
    const variantTheme = themeMode === "light" ? lightMode : darkMode;
    document.documentElement.style.setProperty(
      "--primaryColor",
      variantTheme.primaryColor
    );
    document.documentElement.style.setProperty(
      "--secondaryColor",
      variantTheme.secondaryColor
    );
    document.documentElement.style.setProperty(
      "--statusError",
      variantTheme.statusError
    );
    document.documentElement.style.setProperty(
      "--statusValidation",
      variantTheme.statusValidation
    );
    document.documentElement.style.setProperty(
      "--defaultWhite",
      variantTheme.defaultWhite
    );
    Object.keys(variantTheme.greyscales).forEach((greyScale) =>
      document.documentElement.style.setProperty(
        "--grey" + greyScale,
        variantTheme.greyscales[
          greyScale as unknown as keyof typeof variantTheme.greyscales
        ]
      )
    );

    document.documentElement.style.setProperty(
      "--greyBackground",
      variantTheme.greyBackground
    );
    (global as any).ee42Callback = setThemeMode;

    return createTheme({
      typography: variantTheme.defaultFontFamily,
      palette: {
        mode: variantTheme.mode,
        error: {
          main: variantTheme.statusError,
        },
        success: {
          main: variantTheme.statusValidation,
        },
        grey: variantTheme.greyscales,
        common: {
          white: variantTheme.defaultWhite,
        },
        background: {
          default: variantTheme.background,
        },
        primary: {
          main: variantTheme.primaryColor,
        },
        secondary: {
          main: variantTheme.secondaryColor,
        },
      },
      components: {
        MuiInputBase: {
          styleOverrides: {
            input: {
              '&[type="number"]::-webkit-outer-spin-button': {
                WebkitAppearance: "none",
              },
              '&[type="number"]::-webkit-inner-spin-button': {
                WebkitAppearance: "none",
              },
            },
          },
        },
      },
    });
  }, [themeMode]);

  return (
    <ThemeProvider theme={theme}>
      <ApplicationToolbar></ApplicationToolbar>
      <div className={styles.MainContainer}>
        <div className={styles.board}>
          <Chessboard id="BasicBoard" />
        </div>
      </div>
    </ThemeProvider>
  );
}

export default App;
