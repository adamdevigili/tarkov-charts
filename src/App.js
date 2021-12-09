import React, { useEffect, useState } from 'react';

import './App.css';

import Plot from 'react-plotly.js';
import createTracesFromJSON from './Traces.js'

import Spinner from 'react-bootstrap/Spinner'
import Container from 'react-bootstrap/Container'
import Col from 'react-bootstrap/Col'
import Row from 'react-bootstrap/Row'

import TCNavbar from './components/Navbar';

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
    const apiUrl = endpoint + "/api/ammo"
    fetch(apiUrl, {
      headers: {
        'X-Tarkov-Charts-API-Key': process.env.REACT_APP_TC_API_KEY
      }
    })
      .then((res) => res.json())
      .then((ammoData) => {
        setAppState({ loading: false, ammoData: ammoData });
      });
  }, [setAppState]);

  let finalTraces
  let updatedAt
  if (appState.ammoData) {
    finalTraces = createTracesFromJSON(appState.ammoData)
    updatedAt = appState.ammoData["_updated_at"]
  }

  const plot = <Plot
    config={{displayModeBar: false}}
    data={finalTraces}
    layout={{
      paper_bgcolor:"rgb(230,230,230)",
      height: 1200,
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
        camera: {
          center: {x: 0.05, y: 0.075, z: -0.15},
          eye: {x: 1.35, y: 1.4, z: 1}
        }
      }
    }}
    useResizeHandler={true}
    style={{width: "100%", height: "100%"}}
    onUpdate={() => console.log(plot)}
  />

  return (
    <div>
      <div>
        <TCNavbar/>
      </div>
      <div style={{backgroundColor: '#262626'}}>
        {appState.loading ? (
          <div>
            <div style={{display: 'flex',  justifyContent:'center', alignItems:'center', height: '100vh', color: '#e6e6e6'}}>
              <div>
                <Spinner animation="border" role="status">
                  <span className="sr-only"></span>
                </Spinner>
              </div>
            </div>
          </div>
          ):(
            <Container fluid className='gx-0'>
            <Row className='gx-0'>
              <Col style={{color: '#e6e6e6'}}>
                <div>Single-click a caliber to remove/add it to the graph</div>
                <div>Double-click a caliber to isolate it</div>
                <div>Left click+drag to rotate</div>
                <div>Right click+drag to pan</div>
                <div>Mouse wheel to zoom</div>
                <div>Ctrl+click to add single calibers to the graph</div>
              </Col>
              <Col xs="auto" sm="auto" md="auto" lg="auto" xl="auto" xxl="auto"> 
                <div className='main-plot' > </div>
                <div>{plot}</div>
              </Col>
              <Col></Col>
            </Row>

            </Container>

          )}
      </div>
    </div>
  );
}

export default App;
