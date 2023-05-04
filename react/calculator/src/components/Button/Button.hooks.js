import { evaluate } from "mathjs";

export function addInput(setState, value) {
  const isSpecial = (t) => ["*", "-", "+", "/", "."].includes(t);

  return setState((text) => {
    let newText;

    if (isSpecial(text.charAt(text.length - 1)) && isSpecial(value)) {
      newText = text.slice(0, -1) + value;
    } else {
      newText = text + value;
    }

    return newText;
  });
}

export function calculateResult(setState) {
  setState((text) => {
    try {
      return String(evaluate(text || "0"));
    } catch (error) {
      return "Error";
    }
  });
}

export function clear(setState) {
  setState("");
}
