import reactLogo from "./assets/logo.svg";
import "./App.css";

function App() {
  return (
    <div className="App">
      <div className="React-logo-container">
        <img src={reactLogo} className="React-logo" alt="logo" />
      </div>
      <div className="click-counter-container" />
    </div>
  );
}

export default App;
