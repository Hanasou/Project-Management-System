import React, { Component } from 'react';
import ListGroup from 'react-bootstrap/ListGroup';

import classes from './Dashboard.module.css';
import Wrapper from '../../hoc/Wrapper';

class Dashboard extends Component {
    render() {
        return(
            <Wrapper>
                <ListGroup className={classes.Dashboard}>
                    <ListGroup.Item action href="#link1">
                        Link 1
                    </ListGroup.Item>
                    <ListGroup.Item action href="#link2">
                        Link 2
                    </ListGroup.Item>
                    <ListGroup.Item action href="#link3">
                        This one is a button
                    </ListGroup.Item>
                </ListGroup>
            </Wrapper>
        )
    }
}

export default Dashboard;