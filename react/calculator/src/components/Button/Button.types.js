import { clear, calculateResult, addInput } from "./Button.hooks";

// yes, I know, this is a little over creative lol.
const buttonTypes = Object.freeze({
  Number: Symbol("the button type is a number"),
  Operator: Symbol("the button type is an operator."),
  Clear: Symbol("the button type is to clear the screen."),
  Result: Symbol("the button type is to calculate the result."),
});

export const buttonTypesConfig = new Map();
buttonTypesConfig.set(buttonTypes.Number, { className: "", handler: addInput });
buttonTypesConfig.set(buttonTypes.Operator, {
  className: "operator",
  handler: addInput,
});
buttonTypesConfig.set(buttonTypes.Clear, {
  className: "clear",
  handler: clear,
});
buttonTypesConfig.set(buttonTypes.Result, {
  className: "",
  handler: calculateResult,
});

export default buttonTypes;
