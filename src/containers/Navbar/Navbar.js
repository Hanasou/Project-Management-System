import React, { Component } from 'react';
import { connect } from 'react-redux';
import Form from 'react-bootstrap/Form';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';

import Wrapper from '../../hoc/Wrapper';
import Button from '../../components/UI/Button/Button';
import Input from '../../components/UI/Input/Input';
import * as actions from '../../store/actions/index';

class CustomNavbar extends Component {

    state = {
        controls: {
            email: "",
            password: ""
        }
    };

    // Change state whenever inputs change.
    inputChangedHandler = (event, inputId) => {
        const updatedControls = {
            ...this.state.controls,
            [inputId]: event.target.value
        };
        this.setState({controls: updatedControls});
    };

    // Submit login form.
    loginHandler = (event) => {
        event.preventDefault();
        this.props.onAuthLogin(this.state.controls.email, this.state.controls.password);
    }

    // Log user out
    logoutHandler = (event) => {
        event.preventDefault();
        this.props.onAuthLogout();
    }

    render() {
        let content = null
    
        if (this.props.type === "auth") {
            content = (
                <Wrapper>
                    <Navbar.Brand>Sign In</Navbar.Brand>
                    <Form inline onSubmit={this.loginHandler}>
                        <Form.Group controlId="formLoginEmail" style={{"marginRight": "10px"}}>
                            <Input inputType="email" 
                            inputPlaceholder="Enter email"
                            changed={(event) => this.inputChangedHandler(event, "email")}/>
                        </Form.Group>
                        <Form.Group controlId="formLoginPassword" style={{"marginRight": "10px"}}>
                            <Input inputType="password" 
                            inputPlaceholder="Enter password"
                            changed={(event) => this.inputChangedHandler(event, "password")}/>
                        </Form.Group>
                        <Button btnType="success" type="submit">Sign In</Button>
                    </Form>
                </Wrapper>
            );
        }
        else {
            content = (
                <Wrapper>
                    <Navbar.Brand href="/">Home</Navbar.Brand>
                    <Navbar.Collapse id="basic-navbar-nav">
                        <Nav className="mr-auto">
                        <Nav.Link href="/projects">Projects</Nav.Link>
                        </Nav>
                    </Navbar.Collapse>
                    <Button btnType="success" clicked={this.logoutHandler}>Sign Out</Button>
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

const mapStateToProps = state => {
    return {
        error: state.auth.error,
        isAuthenticated: state.auth.token !== null,
        authRedirectPath: state.auth.authRedirectPath
    };
};

const mapDispatchToProps = dispatch => {
    return {
        onAuthLogin: ( email, password ) => dispatch( actions.authLogin( email, password ) ),
        onAuthLogout: () => dispatch( actions.authLogout() ),
        onSetAuthRedirectPath: () => dispatch(actions.setAuthRedirectPath('/'))
    };
};

export default connect(mapStateToProps, mapDispatchToProps) (CustomNavbar);