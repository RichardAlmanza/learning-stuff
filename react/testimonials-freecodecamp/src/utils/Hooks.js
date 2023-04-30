import { useState } from "react";
import defaultImage from "../assets/images/testimonial-emma.png";

function LoadImage({ image }) {
  const [loaded, setLoaded] = useState(false);
  const srcImage = import(`../assets/images/${image}`).then((module) => {
    setLoaded(true);
    return module.default;
  });

  return [loaded, srcImage, defaultImage];
}

export default LoadImage;
