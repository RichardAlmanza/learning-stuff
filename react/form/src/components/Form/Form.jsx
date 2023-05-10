/* eslint-disable jsx-a11y/label-has-associated-control */
import { useState } from "react";
import "./Form.css";
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
      <FieldSelector
        name="Gender"
        options={genderOptions}
        setState={logChange(setGender)}
        state={gender}
      />
      <label htmlFor="description">Description</label>
      <textarea
        onChange={(event) => {
          logChange(setDescription)(event.target.value);
        }}
        id="description"
        maxLength="3000"
      />
      <button type="submit">Submit</button>
    </form>
  );
}

export default Form;
