import React from 'react';
import Card from 'react-bootstrap/Card';

import classes from './Card.module.css';
import Button from '../Button/Button';

const card = (props) => {
    return(
        <Card className={classes.Card}>
            <Card.Body>
                <Card.Title>{props.title}</Card.Title>
                <Card.Text>
                    {props.desc}
                </Card.Text>
                <Button btnType="primary">View</Button>
            </Card.Body>
        </Card>
    );
};

export default card;