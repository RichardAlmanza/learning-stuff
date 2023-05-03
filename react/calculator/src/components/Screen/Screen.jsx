import PropTypes from "prop-types";
import "./Screen.css";

function Screen({ input }) {
  return <div className="input">{input}</div>;
}

Screen.propTypes = {
  input: PropTypes.number,
};

Screen.defaultProps = {
  input: "",
};

export default Screen;
