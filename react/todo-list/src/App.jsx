import "./App.css";
import logo from "./logo.svg";
import TodoList from "./components/TodoList/TodoList";

function App() {
  return (
    <div className="app">
      <div className="app-logo-container">
        <img className="app-logo" src={logo} alt="logo" />
      </div>
      <TodoList />
    </div>
  );
}

export default App;
