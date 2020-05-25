import React from 'react';
import Button from 'react-bootstrap/button';


const button = (props) => {
    return(
        <Button 
            variant={props.btnType}
            type={props.type}
            onClick={props.clicked}
            >{props.children}</Button>
    )
}

export default button;