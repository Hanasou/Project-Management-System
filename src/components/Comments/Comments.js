import React, { Component } from 'react';
import { connect } from 'react-redux';

import classes from './Comments.module.css';
import Wrapper from '../../hoc/Wrapper';
import Button from '../UI/Button/Button';
import Input from '../UI/Input/Input';
import Form from 'react-bootstrap/Form';
import Comment from './Comment/Comment';
import * as actions from '../../store/actions/index';

class Comments extends Component {
    state = {
        comment: "",
        showButton: false
    }

    inputChangedHandler = (event) => {
        if (this.state.comment === "") {
            this.setState({showButton: false})
        }

        this.setState({
            comment: event.target.value,
            showButton: true
        });
    }

    submitCommentHandler = (event) => {
        event.preventDefault();
        console.log("Submit comment called");
        const newComment = {
            IssueID: this.props.issue.IssueID,
            Email: this.props.email,
            Content: this.state.comment
        }
        this.props.onAddComment(this.props.token, newComment);
    }
    
    render() {
        let button = null;
        if (this.state.showButton === true) {
            button = (
                <Button
                    btnType="primary"
                    type="submit">
                        Comment
                </Button>
            )
        } else {
            button = null;
        }

        let comments = this.props.comments.map(comment => (
            <Comment
                key={comment.CommentID} 
                email={comment.Email}>
                    {comment.Content}
            </Comment>
        ))

        return (
            <Wrapper>
                <h6>Comments</h6>
                <Form onSubmit={this.submitCommentHandler}>
                    <Form.Group controlId="formComment">
                        <Input 
                            inputType="text"
                            inputAs="textarea"
                            rows="3"
                            inputPlaceholder="Add a comment"
                            changed={(event) => this.inputChangedHandler(event)}/>
                        {button}
                    </Form.Group>
                </Form>
                <ul>
                    {comments}
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
        onAddComment: (token, commentData) => dispatch(actions.addComment(token, commentData))
    }
}

export default connect(mapStateToProps, mapDispatchToProps) (Comments);