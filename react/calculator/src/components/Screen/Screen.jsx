import PropTypes from "prop-types";
import "./Screen.css";

function Screen({ input }) {
  return <div className="input">{input}</div>;
}

Screen.propTypes = {
  input: PropTypes.string.isRequired,
};

export default Screen;
