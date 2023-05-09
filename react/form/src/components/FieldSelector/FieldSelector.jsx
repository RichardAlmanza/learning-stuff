import PropTypes from "prop-types";

function FieldSelector({ name, options, setState, state }) {
  const forId = name.toLowerCase();

  const rOptions = options.map((opt) => {
    const vOpt = opt.toLowerCase();
    const id = `${name}${opt}`;

    return (
      <option key={id} value={vOpt}>
        {opt}
      </option>
    );
  });

  const handler = (event) => {
    setState(event.target.value);
  };

  return (
    <label htmlFor={forId}>
      {name}:
      <select name={name} id={forId} onChange={handler} value={state}>
        <option value="">Select {name}</option>
        {rOptions}
      </select>
    </label>
  );
}

FieldSelector.propTypes = {
  name: PropTypes.string.isRequired,
  options: PropTypes.arrayOf(PropTypes.string).isRequired,
  setState: PropTypes.func.isRequired,
  state: PropTypes.string.isRequired,
};

export default FieldSelector;
