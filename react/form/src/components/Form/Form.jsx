/* eslint-disable jsx-a11y/label-has-associated-control */
import { useState } from "react";
import "./Form.css";
import FieldSelector from "../FieldSelector/FieldSelector";
import FieldText from "../FieldText/FieldText";
import logFunc from "../../utils/log";

function Form() {
  const genderOptions = ["Male", "Female", "Other"];

  const [gender, setGender] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [description, setDescription] = useState("");

  const handlerEvent = (f) => {
    return (event) => f(event.target.value);
  };

  const handleDescription = handlerEvent(logFunc(setDescription));
  const handleEmail = handlerEvent(logFunc(setEmail));
  const handlePassword = handlerEvent(logFunc(setPassword));
  const handleGender = handlerEvent(logFunc(setGender));

  const handleSubmit = (event) => {
    event.preventDefault();
    const a = { email, password, gender, description };
    console.log(a);
  };

  return (
    <form onSubmit={handleSubmit}>
      <FieldText
        name="Email"
        type="email"
        setState={handleEmail}
        state={email}
      />
      <FieldText
        name="Password"
        type="password"
        setState={handlePassword}
        state={password}
      />
      <FieldSelector
        name="Gender"
        options={genderOptions}
        setState={handleGender}
        state={gender}
      />
      <label htmlFor="description">Description</label>
      <textarea
        onChange={handleDescription}
        id="description"
        maxLength="3000"
      />
      <button type="submit">Submit</button>
    </form>
  );
}

export default Form;
