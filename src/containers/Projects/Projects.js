import React, { Component } from 'react';
import Card from 'react-bootstrap/Card';
import Button from '../../components/UI/Button/Button';

import Wrapper from '../../hoc/Wrapper';
import classes from './Projects.module.css';
import Dashboard from '../../components/Dashboard/Dashboard';

class Projects extends Component {
    render() {
        return(
            <Wrapper>
                <Dashboard />
                <div className={classes.Projects}>
                <Card style={{ width: '18rem' }}>
                    <Card.Body>
                        <Card.Title>Card Title</Card.Title>
                        <Card.Text>
                            Some quick example text to build on the card title and make up the bulk of
                            the card's content.
                        </Card.Text>
                        <Button btnType="primary">Go somewhere</Button>
                    </Card.Body>
                </Card>
                </div>
            </Wrapper>
        );
    };
};

export default Projects;