import "./Calculator.css";
import Button from "../Button/Button";
import buttonTypes from "../Button/Button.types";
import Screen from "../Screen/Screen";

function Calculator() {
  return (
    <div className="calculator-container">
      <div className="row">
        <Screen input={0.1} />
      </div>
      <div className="row">
        <Button value={1} />
        <Button value={2} />
        <Button value={3} />
        <Button value="+" type={buttonTypes.Operator} />
      </div>
      <div className="row">
        <Button value={4} />
        <Button value={5} />
        <Button value={6} />
        <Button value="-" type={buttonTypes.Operator} />
      </div>
      <div className="row">
        <Button value={7} />
        <Button value={8} />
        <Button value={9} />
        <Button value="*" type={buttonTypes.Operator} />
      </div>
      <div className="row">
        <Button value="=" />
        <Button value={0} />
        <Button value="." />
        <Button value="/" type={buttonTypes.Operator} />
      </div>
      <div className="row">
        <Button value="Clear" type={buttonTypes.Clear} />
      </div>
    </div>
  );
}

export default Calculator;
