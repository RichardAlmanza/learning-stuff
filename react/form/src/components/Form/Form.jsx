import { useState } from "react";
import FieldSelector from "../FieldSelector/FieldSelector";
import FieldText from "../FieldText/FieldText";

function Form() {
  const genderOptions = ["Male", "Female", "Other"];
  const [gender, setGender] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [description, setDescription] = useState("");

  const logChange = (setter) => {
    return (val) => {
      console.log(val);
      setter(val);
    };
  };

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();

        const a = { email, password, gender, description };

        console.log(a);
      }}
    >
      <FieldText
        name="Email"
        type="email"
        setState={logChange(setEmail)}
        state={email}
      />
      <FieldText
        name="Password"
        type="password"
        setState={logChange(setPassword)}
        state={password}
      />
      <FieldText
        name="Description"
        type="text"
        setState={logChange(setDescription)}
        state={description}
      />
      <FieldSelector
        name="Gender"
        options={genderOptions}
        setState={logChange(setGender)}
        state={gender}
      />
      <button type="submit">Submit</button>
    </form>
  );
}

export default Form;
