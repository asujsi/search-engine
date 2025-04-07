import React from 'react';

const ResultsList = ({ results, time, count }) => {
  return (
    <div>
      <div className="text-sm text-gray-600 mb-2">
        {count} result{count !== 1 && 's'} found in {time} ms
      </div>
      {results.map((item, index) => (
        <div key={index} className="p-3 mb-2 border rounded">
          <p className="font-bold text-lg">{item.Message}</p>
          <p className="text-sm text-gray-700">{item.MessageRaw}</p>
          <p className="text-xs text-gray-500">Sender: {item.Sender}</p>
        </div>
      ))}
    </div>
  );
};

export default ResultsList;
