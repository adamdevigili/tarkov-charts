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
  let releaseVersion = process.env.REACT_APP_RELEASE_VERSION
  if (appState.ammoData) {
    finalTraces = createTracesFromJSON(appState.ammoData)
    updatedAt = appState.ammoData["_updated_at"]
  }

  const plot = <Plot
    config={{displayModeBar: false}}
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
        camera: {
          center: {x: 0.05, y: 0.075, z: -0.15},
          eye: {x: 1.35, y: 1.4, z: 1}
        }
      }
    }}
    useResizeHandler={true}
    style={{width: "100%", height: "100%"}}
    // onUpdate={() => console.log(plot)}
  />

  return (
    <div style={{backgroundColor: '#262626', height: '100vh'}}>
      <div>
        <TCNavbar/>
      </div>
      <div>
        {appState.loading ? (
          <div style={{display: 'flex',  justifyContent:'center', alignItems:'center', height: '100vh'}}>
            <div>
              <Spinner animation="border" role="status" style={{color: '#e6e6e6'}}/>
            </div>
          </div>
          ):(
          <Container fluid className='gx-0'>
            <Row className='gx-0'>
              <Col style={{color: '#e6e6e6'}}>
                <ul>
                  <li><b>Single-click</b> a caliber to remove/add it to the graph</li>
                  <li><b>Double-click</b> a caliber to isolate it</li>
                  <li><b>Left click+drag</b> to rotate</li>
                  <li><b>Right click+drag</b> to pan</li>
                  <li><b>Mouse wheel</b> to zoom</li>
                  <li><b>Ctrl+click</b> to add single calibers to the graph</li>
                </ul>
              </Col>
              <Col xs="auto" sm="auto" md="auto" lg="auto" xl="auto" xxl="auto"> 
                <div className='main-plot'> </div>
                <div>{plot}</div>
              </Col>
              <Col className='justify-content-end'>
                <div className="release-version">
                <a href={"https://github.com/adamdevigili/tarkov-charts/releases/tag/" + releaseVersion} rel="noreferrer">
                  {releaseVersion}
                </a>
                </div>
              </Col>
            </Row>
          </Container>
          )}
      </div>
    </div>
  );
}

export default App;
