import React, { useState } from 'react';
import SearchBar from './components/SearchBar';
import ResultsList from './components/ResultsList';
import axios from 'axios';

const App = () => {
  const [results, setResults] = useState([]);
  const [count, setCount] = useState(0);
  const [time, setTime] = useState(0);
  const [loading, setLoading] = useState(false);

  const handleSearch = async (query) => {
    setLoading(true);
    try {
      const res = await axios.get(`http://localhost:8080/search?query=${encodeURIComponent(query)}`);
      setResults(res.data.results);
      setCount(res.data.matches);
      setTime(res.data.time_ms);
    } catch (err) {
      console.error("Search error:", err);
      setResults([]);
      setCount(0);
      setTime(0);
    }
    setLoading(false);
  };

  return (
    <div className="max-w-3xl mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">Log Search Engine</h1>
      <SearchBar onSearch={handleSearch} />
      {loading ? <p>Loading...</p> : <ResultsList results={results} count={count} time={time} />}
    </div>
  );
};

export default App;
