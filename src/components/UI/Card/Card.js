import React from 'react';
import Card from 'react-bootstrap/Card';
import Nav from 'react-bootstrap/Nav';

import classes from './Card.module.css';
import Wrapper from '../../../hoc/Wrapper';
//import Button from '../../UI/Button/Button';

const card = (props) => {
    const cardType = props.cardType
    let contents = null;
    if (cardType === "Project") {
        contents = (
            <Wrapper>
                <Card className={classes.ProjectCard}>
                    <Card.Body>
                        <Card.Title>{props.title}</Card.Title>
                        <Card.Text>
                            {props.desc}
                        </Card.Text>
                        <Nav.Link
                            href={props.path+"/"+props.id}>View</Nav.Link>
                    </Card.Body>
                </Card>
            </Wrapper>
        )
    }
    return(
        <Wrapper>
            {contents}
        </Wrapper>
    );
};

export default card;