import PropTypes from "prop-types";

function FieldText({ name, type, setState, state }) {
  const forId = name.toLowerCase();

  return (
    <>
      <label htmlFor={forId}>{name}</label>
      <input id={forId} type={type} onChange={setState} value={state} />
    </>
  );
}

FieldText.propTypes = {
  name: PropTypes.string.isRequired,
  type: PropTypes.string.isRequired,
  setState: PropTypes.func.isRequired,
  state: PropTypes.string.isRequired,
};

export default FieldText;
