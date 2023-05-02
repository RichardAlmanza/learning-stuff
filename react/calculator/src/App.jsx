import logo from "./logo.svg";
import "./App.css";
import Calculator from "./components/Calculator/Calculator";

function App() {
  return (
    <div className="App">
      <div className="app-logo-container">
        <img src={logo} className="app-logo" alt="logo" />
      </div>
      <Calculator />
    </div>
  );
}

export default App;
