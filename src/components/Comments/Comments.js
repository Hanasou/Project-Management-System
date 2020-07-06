import React, { Component } from 'react';
import { connect } from 'react-redux';

import classes from './Comments.module.css';
import Wrapper from '../../hoc/Wrapper';
import Input from '../UI/Input/Input';
import Form from 'react-bootstrap/Form';
import Comment from './Comment/Comment';
import * as actions from '../../store/actions/index';

class Comments extends Component {
    state = {
        comment: ""
    }

    inputChangedHandler = (event) => {
        this.setState({
            comment: event.target.value
        });
    }
    
    render() {
        return (
            <Wrapper>
                <h6>Comments</h6>
                <Form.Group controlId="formComment">
                    <Input 
                        inputType="text"
                        inputAs="textarea"
                        rows="3"
                        inputPlaceholder="Add a comment"
                        changed={(event) => this.inputChangedHandler(event)}/>
                </Form.Group>
                <ul>
                    <Comment email="email">Comment Body 1</Comment>
                    <Comment email="email2">Comment Body 2</Comment>
                </ul>
            </Wrapper>
        )
    }
}

const mapStateToProps = state => {
    return {
        token: state.auth.token,
        email: state.auth.email,
        comments: state.comments.comments,
        issue: state.issues.issue
    };
};

const mapDispatchToProps = dispatch => {
    return {
        onAddComment: (token, commentData) => dispatch(actions.addComment(token, commentData)),
        onGetComments: (token, issueID) => dispatch(actions.getComments(token, issueID))
    }
}

export default connect(mapStateToProps, mapDispatchToProps) (Comments);