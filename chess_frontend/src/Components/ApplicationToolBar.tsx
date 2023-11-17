import { Toolbar, AppBar, Button, ButtonBase } from "@mui/material";
import styles from "./ApplicationToolbar.module.css";
import MenuButton from "./MenuButton";

export default function ApplicationToolbar() {
  return (
    <AppBar className={styles.appBar}>
      <Toolbar className={styles.toolbar}>
        <MenuButton text="ChEss"></MenuButton>
        <div className={styles.menuItemsContainer}>
          <MenuButton text="Account"></MenuButton>
          <MenuButton text="Play"></MenuButton>
        </div>
      </Toolbar>
    </AppBar>
  );
}
