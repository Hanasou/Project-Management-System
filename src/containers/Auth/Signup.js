import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Redirect } from 'react-router-dom';
import Form from 'react-bootstrap/Form';

import classes from './Signup.module.css';
import Wrapper from '../../hoc/Wrapper';
import Button from '../../components/UI/Button/Button';
import Input from '../../components/UI/Input/Input';
import * as actions from '../../store/actions/index';

class Signup extends Component {
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

    // Submit signup form
    submitHandler = event => {
        event.preventDefault()
        this.props.onAuthSignup( this.state.controls.email, this.state.controls.password);
    }

    render() {
        let authRedirect = null;
        if (this.props.isAuthenticated) {
            authRedirect = <Redirect to={this.props.authRedirectPath}/>
        }
        return(
            <Wrapper>
                {authRedirect}
                <Form className={classes.Signup} onSubmit={this.submitHandler}>
                    <h1>Sign Up</h1>
                    <Form.Group controlId="formSignupEmail">
                        <Form.Label>Email address</Form.Label>
                        <Input 
                            inputType="email" 
                            inputPlaceholder="Enter email"
                            changed={(event) => this.inputChangedHandler(event, "email")}/>
                    </Form.Group>

                    <Form.Group controlId="formSignupPassword">
                        <Form.Label>Password</Form.Label>
                        <Input 
                            inputType="password" 
                            inputPlaceholder="Enter password"
                            changed={(event) => this.inputChangedHandler(event, "password")}/>
                    </Form.Group>
                    <Button btnType="success" type="submit">Sign Up</Button>
                </Form>
            </Wrapper>
        );
    };
};

const mapStateToProps = state => {
    return {
        error: state.auth.error,
        isAuthenticated: state.auth.token !== null,
        authRedirectPath: state.auth.authRedirectPath
    };
};

const mapDispatchToProps = dispatch => {
    return {
        onAuthSignup: ( email, password ) => dispatch( actions.authSignup( email, password ) ),
        onSetAuthRedirectPath: () => dispatch(actions.setAuthRedirectPath('/'))
    };
};

export default connect(mapStateToProps, mapDispatchToProps) (Signup);