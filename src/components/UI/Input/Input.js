import React from 'react';
import Form from 'react-bootstrap/Form';

const input = (props) => {
    return(
        <Form.Control 
            type={props.inputType}
            as={props.inputAs}
            rows={props.rows} 
            placeholder={props.inputPlaceholder} 
            onChange={props.changed}
            readOnly={props.readOnly}
            plaintext={props.plaintext}
            defaultValue={props.children}/>
    )
}

export default input;