import PropTypes from "prop-types";
import "./Testimonial.css";

const EmmaImage = require("../../images/testimonial-emma.png");

function Testimonial({ name, role, country, company, testimonial }) {
  return (
    <div className="testimonial-container">
      <img className="testimonial-image" src={EmmaImage} alt="emma" />
      <div className="testimonial-text-container">
        <p className="testimonial-name">
          {name} in {country}
        </p>
        <p className="testimonial-role">
          {role} at {company}
        </p>
        <p className="testimonial-text">{testimonial}</p>
      </div>
    </div>
  );
}

Testimonial.propTypes = {
  name: PropTypes.string.isRequired,
  role: PropTypes.string.isRequired,
  country: PropTypes.string.isRequired,
  company: PropTypes.string.isRequired,
  testimonial: PropTypes.string.isRequired,
};

export default Testimonial;
