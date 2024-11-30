import React from "react";

function Cache({ cache }) {
  return (
    <div className="cache-container">
      <h2>Cache Items:</h2>
      <div className="cache-list">
        {cache.map((word, index) => (
          <div className="cache-item" key={index}>
            {word}
          </div>
        ))}
      </div>
    </div>
  );
}

export default Cache;
