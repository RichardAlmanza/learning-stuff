// yes, I know, this is a little over creative lol.
const buttonTypes = Object.freeze({
  Number: Symbol("the button type is a number"),
  Operator: Symbol("the button type is an operator."),
  Clear: Symbol("the button type is an operator."),
});

export const buttonTypesConfig = new Map();
buttonTypesConfig.set(buttonTypes.Number, "");
buttonTypesConfig.set(buttonTypes.Operator, "operator");
buttonTypesConfig.set(buttonTypes.Clear, "clear");

export default buttonTypes;
