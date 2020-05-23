import React, { Component } from 'react';
import Form from 'react-bootstrap/Form';
import Navbar from 'react-bootstrap/Navbar';

import Wrapper from '../../hoc/Wrapper';
import Button from '../../components/UI/Button/Button';
import Input from '../../components/UI/Input/Input';

class CustomNavbar extends Component {
    render() {
        let content = null
    
        if (this.props.type === "auth") {
            content = (
                <Wrapper>
                    <Navbar.Brand href="#home">Sign In</Navbar.Brand>
                    <Form inline>
                        <Form.Group controlId="formLoginEmail" style={{"marginRight": "10px"}}>
                            <Input inputType="email" inputPlaceholder="Enter email"/>
                        </Form.Group>
                        <Form.Group controlId="formLoginPassword" style={{"marginRight": "10px"}}>
                            <Input inputType="password" inputPlaceholder="Enter password"/>
                        </Form.Group>
                        <Button btnType="success">Sign In</Button>
                    </Form>
                </Wrapper>
            );
        }
        else {
            content = (
                <Wrapper>
                    <Navbar.Brand href="#home">Welcome</Navbar.Brand>
                    <Button btnType="success">Sign Out</Button>
                </Wrapper>
            );
        }

        return(
            <Navbar className="bg-light justify-content-between">
                {content}
            </Navbar>
        )
    }

}

export default CustomNavbar;