import PropTypes from "prop-types";

function FieldText({ name, type, setState, state }) {
  const forId = name.toLowerCase();

  const handler = (event) => {
    setState(event.target.value);
  };

  return (
    <label htmlFor={forId}>
      {name}:
      <input type={type} onChange={handler} value={state} />
    </label>
  );
}

FieldText.propTypes = {
  name: PropTypes.string.isRequired,
  type: PropTypes.string.isRequired,
  setState: PropTypes.func.isRequired,
  state: PropTypes.string.isRequired,
};

export default FieldText;
