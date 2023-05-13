export default function logFunc(setter) {
  return (val) => {
    console.log(val);
    setter(val);
  };
}
