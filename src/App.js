import React from 'react';
// import { useEffect, useState } from 'react';
import './App.css';
import Plot from 'react-plotly.js';
// import rawData from './data/ammo.json';

function App() {
  // const traces = Array(1).fill(0).map((_, i) => {
  //   const {index, arr} = randomValues(20, 3);
  //   return {
  //     x: Array(20).fill(i),
  //     y: index,
  //     z: arr,
  //     type: 'scatter3d',
  //     mode: 'lines',
  //   }
  // });

  const traces = [
    {      
      type: 'scatter3d',
      mode: 'lines+markers',
      name: '5.56x45 mm',
      x:[85,50,45,40], // DAMAGE
      y:[3,28,43,53], // PEN
      z:[198,178,2784,2988], // PRICE
      text: ["Warmage", "M855", "M855A1", "M995"],
      marker: {color: 'red', size: 5}
    },
    {      
      type: 'scatter3d',
      mode: 'lines+markers',
      name: '7.62x51 mm',
      x:[107,80,70,67], // DAMAGE
      y:[15,41,64,70], // PEN
      z:[261,377,2129,3709], // PRICE
      text: ["Ultra Nosler", "M80", "M61", "M993"],
      marker: {color: 'blue', size: 5}
    },
  ]

  console.log(traces)

  return (
    <Plot
      data={traces}
      layout={{
        height: 1200,
        width: 1200,
        title: `Tarko Ammo: Damage/Penetration/Price`,
        autosize: true,
        scene: {
          xaxis: {
            title: "damage"
          },
          yaxis: {
            title: "penetration"
          },
          zaxis: {
            title: "cost (₽)"
          }
        }
      }}
      useResizeHandler={true}
      style={{width: "100%", height: "100%"}}

    />
  );
}

export default App;
