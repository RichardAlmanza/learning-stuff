import "./Button.css";
import PropTypes from "prop-types";
import { evaluate } from "mathjs";
import buttonTypes, { buttonTypesConfig } from "./Button.types";

function getAction(value, setState, type) {
  if (type === buttonTypes.Number || type === buttonTypes.Operator) {
    return () => {
      setState((text) => text + value);
    };
  }

  if (type === buttonTypes.Result) {
    return () => {
      setState((text) => evaluate(text || "0"));
    };
  }

  if (type === buttonTypes.Clear) {
    return () => {
      setState("");
    };
  }

  return () => console.log("something got wrong");
}

function Button({ value, setState, type }) {
  const bType = buttonTypesConfig.get(type);
  const className = `button-container ${bType}`.trimEnd();

  const clickHandler = getAction(value, setState, type);

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
