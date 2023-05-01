import PropTypes from "prop-types";
import Testimonial from "../Testimonial/Testimonial";

// eslint-disable-next-line react/prop-types
function Testimonials({ list }) {
  // eslint-disable-next-line react/prop-types
  const testimonials = list.map((t) => (
    <Testimonial
      key={t.id}
      name={t.name}
      role={t.role}
      country={t.country}
      testimonial={t.testimonial}
      company={t.company}
      image={t.image}
    />
  ));

  return <> {testimonials}</>;
}

Testimonials.protoTypes = {
  // eslint-disable-next-line react/forbid-foreign-prop-types
  list: PropTypes.arrayOf(PropTypes.objectOf(Testimonial.propTypes)).isRequired,
};

export default Testimonials;
