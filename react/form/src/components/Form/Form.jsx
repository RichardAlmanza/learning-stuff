import { useState } from "react";
import FieldSelector from "../FieldSelector/FieldSelector";
import FieldText from "../FieldText/FieldText";

function Form() {
  const genderOptions = ["Male", "Female", "Other"];
  const [gender, setGender] = useState("");
  const [email, setEmail] = useState("");

  const logChange = (setter) => {
    return (val) => {
      console.log(val);
      setter(val);
    };
  };

  return (
    <form>
      <FieldText
        name="Email"
        type="email"
        setState={logChange(setEmail)}
        state={email}
      />
      <FieldSelector
        name="Gender"
        options={genderOptions}
        setState={logChange(setGender)}
        state={gender}
      />
    </form>
  );
}

Form.propTypes = {};

export default Form;
