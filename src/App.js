import React, { useEffect, useState } from 'react';
import './App.css';

import Plot from 'react-plotly.js';

import SplitPane from 'react-split-pane';

import { Button } from 'react-bootstrap';
import Spinner from 'react-bootstrap/Spinner'
import Container from 'react-bootstrap/Container'
import Menu from './components/Menu';
import { Col, Row, OverlayTrigger, Tooltip } from 'react-bootstrap';

import createTracesFromJSON from './Traces.js'

import rawAmmoData from './data/get_ammo.json';

function App() {
  const [appState, setAppState] = useState({
    loading: false,
    ammoData: null,
  });

  useEffect(() => {
    setAppState({ loading: true });
    let endpoint = ""
    if (process.env.REACT_APP_API_ENDPOINT) {
      endpoint = process.env.REACT_APP_API_ENDPOINT
    } else {
      endpoint = "https://www.tarkov-charts.com"
    }
    const apiUrl = endpoint + "/api/get/ammo"

    if (process.env.NODE_ENV === "development") {
      console.log("development mode, using hardcoded values")
      setAppState({ loading: false, ammoData: rawAmmoData })
    } else {
      fetch(apiUrl, {
        headers: {
          'X-Tarkov-Charts-API-Key': process.env.REACT_APP_TC_API_KEY
        }
      })
      .then((res) => res.json())
      .then((ammoData) => {
        setAppState({ loading: false, ammoData: ammoData });
      });
    }
  }, [setAppState]);

  let finalTraces
  let updatedAt
  if (appState.ammoData) {
    finalTraces = createTracesFromJSON(appState.ammoData)
    updatedAt = appState.ammoData["_updated_at"]
  }

  const plot = <Plot
    config={{displayModeBar: false, responsive: true}}
    data={finalTraces}
    layout={{
      paper_bgcolor:"rgb(230,230,230)",
      height: 1000,
      width: 1200,
      title: `Tarkov Ammo by Caliber: Damage/Penetration/Price (last updated ` + updatedAt + `)`,
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
          title: "Cost (₽, 24hr avg.)",
          range: [0, 5000],
        },
        // bgcolor: "#FFFFFF",
        camera: {
          center: {
            z: -.15
          },
          eye: {
              x: 1.5,
              y: 1.5,
              z: 1,
          }
        }
      
      }
    }}
    useResizeHandler={true}
    // style={{width: "100%", height: "100%"}}
  />


  return (
    <div>
      <Menu/>
      {appState.loading ? (
          <div style={{display: 'flex',  justifyContent:'center', alignItems:'center', height: '100vh'}}>
            <div>          
              <Spinner animation="border" role="status">
                <span className="sr-only"></span>
              </Spinner>
            </div>
          </div>
      ):(
        <Container fluid>
        <Row className='main-row'>
          <Col className='panel-col'>
            <h1>control-panel</h1>
          </Col>
          <Col className='plot-col'>
            <div className='main-plot'>
              <div>{plot}</div>
              {/* <div>Single-click a caliber to remove/add it to the graph</div>
              <div>Double-click a caliber to isolate it</div>
              <div>Left click+drag to rotate</div>
              <div>Right click+drag to pan</div>
              <div>Mouse wheel to zoom</div>
              <div>Ctrl+click to add single calibers to the graph</div> */}
            </div>
          </Col>
        </Row>

        </Container>
        )} 
    </div>

  );
}

export default App;
