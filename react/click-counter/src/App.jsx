import { useState } from "react";
import reactLogo from "./assets/logo.svg";
import Button from "./components/Button/Button";
import Counter from "./components/Counter/Counter";
import "./App.css";

function App() {
  const [nClick, setNClick] = useState(0);

  const addClick = () => {
    setNClick((clicks) => clicks + 1);
  };

  const resetCounter = () => {
    setNClick(0);
  };

  return (
    <div className="app">
      <div className="react-logo-container">
        <img src={reactLogo} className="React-logo" alt="logo" />
      </div>
      <div className="click-counter-container">
        <Counter clickCounter={nClick} />
        <Button text="Click" eventHandler={addClick} />
        <Button
          text="Reset"
          isClickButton={false}
          eventHandler={resetCounter}
        />
      </div>
    </div>
  );
}

export default App;
