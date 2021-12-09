import Navbar from "react-bootstrap/Navbar";
import Nav from "react-bootstrap/Nav";
import "./Navbar.css";

import React from "react";

function TCNavbar({ releaseVersion }) {
  return (
    <div>
      <Navbar style={{ backgroundColor: "#020202" }}>
        <Navbar.Brand style={{ color: "#9a8866" }} href="/">
          Tarkov Charts
        </Navbar.Brand>
        <Navbar.Collapse className="justify-content-end">
          <Nav.Link
            style={{ color: "#9a8866" }}
            href={
              "https://github.com/adamdevigili/tarkov-charts/releases/tag/" +
              releaseVersion
            }
          >
            {releaseVersion}
          </Nav.Link>
        </Navbar.Collapse>
      </Navbar>
    </div>
  );
}

export default TCNavbar;
