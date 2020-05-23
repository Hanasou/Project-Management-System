import React, { Component } from 'react';
import Form from 'react-bootstrap/Form';

import classes from './Signup.module.css';
import Wrapper from '../../hoc/Wrapper';
import Button from '../../components/UI/Button/Button';
import Input from '../../components/UI/Input/Input';
import CustomNavbar from '../Navbar/Navbar';

class Signup extends Component {
    render() {
        return(
            <Wrapper>
                <CustomNavbar type="auth"/>
                <Form className={classes.Signup}>
                    <h1>Sign Up</h1>
                    <Form.Group controlId="formSignupEmail">
                        <Form.Label>Email address</Form.Label>
                        <Input inputType="email" inputPlaceholder="Enter email"/>
                    </Form.Group>

                    <Form.Group controlId="formSignupPassword">
                        <Form.Label>Password</Form.Label>
                        <Input inputType="password" inputPlaceholder="Enter password"/>
                    </Form.Group>
                    <Button btnType="success">Sign Up</Button>
                </Form>
            </Wrapper>
        )
    }
}

export default Signup;