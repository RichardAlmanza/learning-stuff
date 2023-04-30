import { useState } from "react";

function LoadImage(image) {
  const [srcImage, setSrcImage] = useState({
    name: "placeholder",
    src: "",
  });

  import(`../assets/images/${image}`)
    .then((module) => module.default)
    .then((src) => {
      setSrcImage({ name: image, src });
    });

  return srcImage;
}

export default LoadImage;
