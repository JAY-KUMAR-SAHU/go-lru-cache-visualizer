import React, { useState, useEffect } from "react";
import Cache from "./components/Cache";
import "./App.css";

function App() {
  const [cache, setCache] = useState([]);
  const [word, setWord] = useState("");
  const [cacheSize, setCacheSize] = useState(6);
  const [newSize, setNewSize] = useState(cacheSize);
  const [showWarning, setShowWarning] = useState(false);

  useEffect(() => {
    fetchCache();
  }, []);

  const fetchCache = async () => {
    const response = await fetch("http://localhost:8080/cache");
    const data = await response.json();
    setCache(data.cache);
    setCacheSize(data.capacity);
  };

  const addWord = async () => {
    if (!word) return;
    await fetch(`http://localhost:8080/add/${word}`, { method: "POST" });
    setWord("");
    fetchCache();
  };

  const changeCacheSize = async () => {
    if (newSize < cacheSize) {
      setShowWarning(true);
    } else {
      await fetch(`http://localhost:8080/set-cache-size/${newSize}`, {
        method: "POST",
      });
      fetchCache();
    }
  };

  const handleCacheSizeChange = (e) => {
    const value = parseInt(e.target.value);
    if (!isNaN(value) && value > 0) {
      setNewSize(value);
    } else {
      setNewSize("");
    }
  };

  const confirmSizeChange = async () => {
    setShowWarning(false);
    await fetch(`http://localhost:8080/set-cache-size/${newSize}`, {
      method: "POST",
    });
    fetchCache();
  };

  const cancelSizeChange = () => {
    setShowWarning(false);
  };

  return (
    <div className="container">
      <h1>GO - LRU Cache Visualizer</h1>

      <div className="cache-info">
        <h2>
          Cache State: {cache.length}/{cacheSize}
        </h2>
      </div>

      <div className="cache-size-input">
        <label>Set Cache Size: </label>
        <input
          type="number"
          value={newSize}
          onChange={handleCacheSizeChange}
          min="1"
          max="100"
        />
        <button onClick={changeCacheSize}>Change Size</button>
      </div>

      {showWarning && (
        <div className="warning-dialog">
          <p>
            Changing the cache size will remove extra data if the new size is
            smaller. Are you sure?
          </p>
          <button className="toRed" onClick={confirmSizeChange}>
            Yes
          </button>
          <button className="toGreen" onClick={cancelSizeChange}>
            Cancel
          </button>
        </div>
      )}

      <Cache cache={cache} />
      <input
        type="text"
        value={word}
        onChange={(e) => setWord(e.target.value)}
        placeholder="Add word"
      />
      <button onClick={addWord}>Add Word</button>
    </div>
  );
}

export default App;
