import reactLogo from "./assets/logo.svg";
import Button from "./components/Button/Button";
import Counter from "./components/Counter/Counter";
import "./App.css";

function App() {
  const addClick = () => {
    console.log("click");
  };

  const resetCounter = () => {
    console.log("reset");
  };

  return (
    <div className="app">
      <div className="react-logo-container">
        <img src={reactLogo} className="React-logo" alt="logo" />
      </div>
      <div className="click-counter-container">
        <Counter clickCounter={5} />
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
