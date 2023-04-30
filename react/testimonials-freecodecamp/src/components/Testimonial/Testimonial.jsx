import "./Testimonial.css";
import PropTypes from "prop-types";
import LoadImage from "../../utils/Hooks";

function Testimonial({ name, role, country, company, testimonial, image }) {
  const [isImageReady, srcImage, srcDefaultImage] = LoadImage(
    `testimonial-${image}.png`
  );

  return (
    <div className="testimonial-container">
      <img
        className="testimonial-image"
        src={isImageReady ? srcImage : srcDefaultImage}
        alt={image}
      />
      <div className="testimonial-text-container">
        <p className="testimonial-name">
          <strong>{name}</strong> in {country}
        </p>
        <p className="testimonial-role">
          {role} at <strong>{company}</strong>
        </p>
        <p className="testimonial-text">&quot;{testimonial}&quot;</p>
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
  image: PropTypes.string.isRequired,
};

export default Testimonial;
