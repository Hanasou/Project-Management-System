import React from 'react';
import Dropdown from 'react-bootstrap/Dropdown';
import DropdownButton from 'react-bootstrap/DropdownButton';

const dropdown = (props) => {
    let dropdownItems = props.items.map( item => (
        <Dropdown.Item 
            key={props.items.indexOf(item)}
            eventKey={item}>{item}</Dropdown.Item>
    ))
    return(
        <DropdownButton
            title={props.title}
            onSelect={props.onSelect}>
            {dropdownItems}
        </DropdownButton>
    )
}

export default dropdown;