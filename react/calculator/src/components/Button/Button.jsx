import "./Button.css";
import PropTypes from "prop-types";
import buttonTypes, { buttonTypesConfig } from "./Button.types";

function Button({ value, setState, type }) {
  const bType = buttonTypesConfig.get(type);
  const className = `button-container ${bType.className}`.trimEnd();
  const clickHandler = () => bType.handler(setState, value);

  return (
    <button type="button" onClick={clickHandler} className={className}>
      {value}
    </button>
  );
}

Button.propTypes = {
  value: PropTypes.string.isRequired,
  setState: PropTypes.func.isRequired,
  type: PropTypes.symbol,
};

Button.defaultProps = {
  type: buttonTypes.Number,
};

export default Button;
