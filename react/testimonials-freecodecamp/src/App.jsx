import "./App.css";
import data from "./assets/data/data.json";
import Testimonials from "./components/Testimonials/Testimonials";

function App() {
  return (
    <div className="App">
      <div className="main-container">
        <h1>Here is what our alumni say about freeCodeCamp:</h1>
        <Testimonials list={data.testimonials} />
      </div>
    </div>
  );
}

export default App;
