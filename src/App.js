import React, { useEffect, useState } from 'react';
// import { useEffect, useState } from 'react';
import './App.css';
import Plot from 'react-plotly.js';
import createTracesFromJSON from './Traces.js'
// import rawData from './data/ammo.json';

function App() {
  const [appState, setAppState] = useState({
    loading: false,
    ammoData: null,
  });

  useEffect(() => {
    setAppState({ loading: true });
    const apiUrl = "https://api.jsonbin.io/v3/b/" + process.env.REACT_APP_JSONBIN_BIN_ID;
    fetch(apiUrl, {
      headers: {
        'X-Master-Key': process.env.REACT_APP_JSONBIN_API_KEY
      }
    })
      .then((res) => res.json())
      .then((ammoData) => {
        setAppState({ loading: false, ammoData: ammoData });
      });
  }, [setAppState]);

  console.log(appState.ammoData)

  const traceColorList = ['red', 'blue', 'green', 'purple']
  const markerSymbolList = []

  // appState.ammoData
  let finalTraces
  if (appState.ammoData) {
    finalTraces = createTracesFromJSON(appState.ammoData)
  }
  // const ammoTraceData = JSON.parse(appState.ammoData)

  const plot = <Plot
    config={{displayModeBar: false}}
    data={finalTraces}
    layout={{
      paper_bgcolor:"rgb(240,240,240)",
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

  return (
    <div> 
      {appState.loading ? (
        <div>loading</div> 
        ):(
          plot
        )} 
    </div>
  );
}

export default App;
