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

  const baseTrace = {
      type: 'scatter3d',
      mode: 'lines+markers+text',
      textposition: 'top center',
      hovertemplate:
      '<b><i>%{text}</i></b><br>' +
      'Damage: %{x}<br>' +
      'Pen: %{y}<br>' +
      'Cost: ₽ %{z}<br>'
  }

  const traces = [
    {      
      ...baseTrace,
      name: '5.56x45 mm',
      x:[85,50,45,40], // DAMAGE
      y:[3,28,43,53], // PEN
      z:[198,178,2784,2988], // PRICE
      text: ["Warmage", "M855", "M855A1", "M995"],
      marker: {color: 'red', size: 5},
    },
    { 
      ...baseTrace,     
      name: '7.62x51 mm',
      x:[107,80,70,67], // DAMAGE
      y:[15,41,64,70], // PEN
      z:[261,377,2129,3709], // PRICE
      text: ["Ultra Nosler", "M80", "M61", "M993"],
      marker: {color: 'blue', size: 5},
    },
    { 
      ...baseTrace,     
      name: '9x39 mm',
      x:[68,58,64,60], // DAMAGE
      y:[38,46,50,55], // PEN
      z:[261,660,1146,2240], // PRICE
      text: ["SP-5", "SP-6", "7N9 SPP", "7N12 BP"],
      marker: {color: 'orange', size: 5},
    },
    { 
      ...baseTrace,     
      name: '4.6x30 mm',
      x:[65,45,43,35], // DAMAGE
      y:[18,36,40,53], // PEN
      z:[103,1370,2178,2867], // PRICE
      text: ["Action SX", "Subsonic SX", "FMJ SX", "AP SX"],
      marker: {color: 'purple', size: 5},
    }
  ]

  return (
    <Plot
      config={{displayModeBar: false}}
      data={traces}
      layout={{
        xaxis: {
            title: "damage",
          },
        height: 1200,
        width: 1200,
        title: `Tarkov Ammo: Damage/Penetration/Price`,
        autosize: true,
        scene: {
          xaxis: {
            title: "Damage",
            range: [200, 0],
          },
          yaxis: {
            title: "Penetration",
            range: [80, 0],
          },
          zaxis: {
            title: "Cost (₽)"
          }
        }
      }}
      useResizeHandler={true}
      style={{width: "100%", height: "100%"}}

    />
  );
}

export default App;
