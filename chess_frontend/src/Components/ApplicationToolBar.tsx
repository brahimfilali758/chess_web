import { Toolbar, AppBar, FormControlLabel } from "@mui/material";
import styles from "./ApplicationToolbar.module.css";
import MenuButton from "./MenuButton";
import { MaterialUISwitch } from "./MaterialUISwitch";

export type TApplicationToolBarParams = {
  setTheme: (theme: "light" | "dark") => void;
};

export default function ApplicationToolbar({
  setTheme,
}: TApplicationToolBarParams) {
  return (
    <AppBar className={styles.appBar}>
      <Toolbar className={styles.toolbar}>
        <div className={styles.LeftSide}>
          <FormControlLabel
            label="Light/Dark"
            control={<MaterialUISwitch sx={{ m: 1 }}></MaterialUISwitch>}
            onChange={(event, checked) =>
              checked ? setTheme("dark") : setTheme("light")
            }
          ></FormControlLabel>
          <MenuButton text="GOChess"></MenuButton>
        </div>

        <div className={styles.menuItemsContainer}>
          <MenuButton text="Account"></MenuButton>
          <MenuButton text="Play"></MenuButton>
        </div>
      </Toolbar>
    </AppBar>
  );
}
