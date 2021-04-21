import Navbar from 'react-bootstrap/Navbar'
import Nav from 'react-bootstrap/Nav'

import React  from 'react';

function Menu() {
    return (
        <div>
        <Navbar bg="dark" variant="dark">
            <Navbar.Brand href="#home">Tarkov Charts</Navbar.Brand>
                <Nav className="mr-auto">
                    {/* <Nav.Link href="#ammo">Ammo</Nav.Link> */}
                    {/* <Nav.Link href="#attachments">Attachments</Nav.Link> */}
                </Nav>
        </Navbar>
        </div>
    )
    }

export default Menu