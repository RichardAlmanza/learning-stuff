import logo from "./logo.svg";
import "./App.css";

function App() {
  return (
    <div className="App">
      <div className="app-logo-container">
        <img src={logo} className="app-logo" alt="logo" />
      </div>
      <div className="calculator-container">
        <div className="row" />
        <div className="row" />
        <div className="row" />
        <div className="row" />
        <div className="row" />
        <div className="row" />
        <div className="row" />
        <div className="row" />
      </div>
    </div>
  );
}

export default App;
