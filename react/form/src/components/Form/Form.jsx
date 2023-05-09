import { useState } from "react";
import FieldSelector from "../FieldSelector/FieldSelector";

function Form() {
  const genderOptions = ["Male", "Female", "Other"];
  const [gender, setGender] = useState("");

  return (
    <form>
      <FieldSelector
        name="Gender"
        options={genderOptions}
        setState={setGender}
        state={gender}
      />
    </form>
  );
}

Form.propTypes = {};

export default Form;
