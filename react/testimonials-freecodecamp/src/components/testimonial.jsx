import React from "react";

function Testimonial() {
  return (
    <div className="testimonial-container">
      <img
        className="testimonial-image"
        src="{require(../images/testimonial-emma.png)}"
        alt="emma"
      />
      <div className="testimonial-text-container">
        <p className="testimonial-name">Emma Bostian in Sweden</p>
        <p className="testimonial-role">Software Engineer at Spotify</p>
        <p className="testimonial-text">"I've always struggled with learning JavaScript. I've taken many courses but freeCodeCamp's course was the one which stuck. Studying JavaScript as well as data structures and algorithms on freeCodeCamp gave me the skills and confidence I needed to land my dream job as a software engineer at Spotify."</p>
      </div>
    </div>
  );
}

export default Testimonial
