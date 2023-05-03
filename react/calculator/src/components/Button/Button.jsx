import "./Button.css";
import PropTypes from "prop-types";

const buttonType = Object.freeze({
  Number: { id: Symbol("the button type is a number"), className: "" },
  Operator: {
    // yes, I know, this is a little over creative lol.
    id: Symbol("the button type is an operator."),
    className: "operator",
  },
});

function getButtonType(element) {
  if (["+", "-", "*", "/"].includes(element)) {
    return buttonType.Operator;
  }

  return buttonType.Number;
}

function Button({ value }) {
  const bType = getButtonType(value);
  const className = `button-container ${bType.className}`.trimEnd();

  return <div className={className}>{value}</div>;
}

Button.propTypes = {
  value: PropTypes.string.isRequired,
};

export default Button;
