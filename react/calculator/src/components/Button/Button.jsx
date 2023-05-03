import "./Button.css";
import PropTypes from "prop-types";
import buttonTypes, { buttonTypesConfig } from "./Button.types";

function Button({ value, type }) {
  const bType = buttonTypesConfig.get(type);
  const className = `button-container ${bType}`.trimEnd();

  return <div className={className}>{value}</div>;
}

Button.propTypes = {
  value: PropTypes.string.isRequired,
  type: PropTypes.symbol,
};

Button.defaultProps = {
  type: buttonTypes.Number,
};

export default Button;
