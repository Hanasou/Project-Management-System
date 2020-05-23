import React from 'react';
import Button from 'react-bootstrap/button';


const button = (props) => {
    return(
        <Button variant={props.btnType}>{props.children}</Button>
    )
}

export default button;