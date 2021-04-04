import React from 'react';
// import { useEffect, useState } from 'react';
import './App.css';
import Plot from 'react-plotly.js';

function randomValues(num, mul) {
  const arr = [];
  const index = [];
  for (let i = 0; i < num; i++) {
    arr.push(Math.random() * mul)
    index.push(i);
  }
  return {index, arr};
}

function App() {
  // const [date, setDate] = useState(null);
  // useEffect(() => {
  //   async function getDate() {
  //     const res = await fetch('/api/date');
  //     const newDate = await res.text();
  //     setDate(newDate);
  //   }
  //   getDate();
  // }, []);

  const traces = Array(3).fill(0).map((_, i) => {
    const {index, arr} = randomValues(20, 3);
    return {
      x: Array(20).fill(i),
      y: index,
      z: arr,
      type: 'scatter3d',
      mode: 'lines'
    }
  });

  return (
    <Plot
      data={traces}
      layout={{
        width: 900,
        height: 800,
        title: `Simple 3D Scatter`
      }}
    />
  );
}

export default App;
