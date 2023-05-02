import PropTypes from "prop-types";
import "./Counter.css";

function Counter({ clickCounter }) {
  return <div className="counter">{clickCounter}</div>;
}

Counter.propTypes = {
  clickCounter: PropTypes.number.isRequired,
};

export default Counter;
