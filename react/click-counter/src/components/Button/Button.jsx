import PropTypes from "prop-types";
import "./Button.css";

function Button({ text, eventHandler, isClickButton }) {
  return (
    <button
      type="button"
      className={isClickButton ? "click-button" : "reset-button"}
      onClick={eventHandler}
    >
      {text}
    </button>
  );
}

Button.propTypes = {
  text: PropTypes.string.isRequired,
  eventHandler: PropTypes.func.isRequired,
  isClickButton: PropTypes.bool,
};

Button.defaultProps = {
  isClickButton: true,
};

export default Button;
