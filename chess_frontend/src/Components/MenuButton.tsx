import { ButtonBase } from "@mui/material";
import styles from "./MenuButton.module.css";

export type TMenuTabParams = {
  text: string;
};

export default function MenuButton({ text }: TMenuTabParams) {
  return (
    <span className={styles.toolbarChild}>
      <ButtonBase className={styles.toolbarButton}>{text}</ButtonBase>
    </span>
  );
}
