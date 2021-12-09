import Navbar from "react-bootstrap/Navbar";
import "./Navbar.css";

import React from "react";

function TCNavbar() {
  return (
    <div>
      <Navbar style={{ backgroundColor: "#020202" }}>
        <Navbar.Brand style={{ color: "#9a8866" }} href="#home">
          Tarkov Charts
        </Navbar.Brand>
        {/* <Nav className="mr-auto">
                        <Nav.Link href="#ammo">Ammo</Nav.Link>
                        <Nav.Link href="#attachments">Attachments</Nav.Link>
                    </Nav> */}
      </Navbar>
    </div>
  );
}

export default TCNavbar;
