import React, { Component } from 'react';
import Nav from 'react-bootstrap/Nav';

import classes from './Dashboard.module.css';
import Wrapper from '../../hoc/Wrapper';

class Dashboard extends Component {
    render() {
        return(
            <Wrapper>
                <Nav defaultActiveKey="/home" className={classes.Dashboard}>
                    <Nav.Link eventKey="link-0">Link</Nav.Link>
                    <Nav.Link eventKey="link-1">Link</Nav.Link>
                    <Nav.Link eventKey="link-2">Link</Nav.Link>
                    <Nav.Link eventKey="disabled" disabled>
                        Disabled
                    </Nav.Link>
                </Nav>
            </Wrapper>
        )
    }
}

export default Dashboard;