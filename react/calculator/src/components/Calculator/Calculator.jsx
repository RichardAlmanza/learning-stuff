import "./Calculator.css";
import { useState } from "react";
import Button from "../Button/Button";
import buttonTypes from "../Button/Button.types";
import Screen from "../Screen/Screen";

function Calculator() {
  const [input, setInput] = useState("");

  return (
    <div className="calculator-container">
      <div className="row">
        <Screen input={input} />
      </div>
      <div className="row">
        <Button value="1" setState={setInput} />
        <Button value="2" setState={setInput} />
        <Button value="3" setState={setInput} />
        <Button value="+" setState={setInput} type={buttonTypes.Operator} />
      </div>
      <div className="row">
        <Button value="4" setState={setInput} />
        <Button value="5" setState={setInput} />
        <Button value="6" setState={setInput} />
        <Button value="-" setState={setInput} type={buttonTypes.Operator} />
      </div>
      <div className="row">
        <Button value="7" setState={setInput} />
        <Button value="8" setState={setInput} />
        <Button value="9" setState={setInput} />
        <Button value="*" setState={setInput} type={buttonTypes.Operator} />
      </div>
      <div className="row">
        <Button value="=" setState={setInput} type={buttonTypes.Result} />
        <Button value="0" setState={setInput} />
        <Button value="." setState={setInput} />
        <Button value="/" setState={setInput} type={buttonTypes.Operator} />
      </div>
      <div className="row">
        <Button value="Clear" setState={setInput} type={buttonTypes.Clear} />
      </div>
    </div>
  );
}

export default Calculator;
