import React, { Component } from 'react';
import { connect } from 'react-redux';
import Modal from 'react-bootstrap/Modal';
import Form from 'react-bootstrap/Form';

import Wrapper from '../../../hoc/Wrapper';
import Button from '../Button/Button';
import Input from '../Input/Input';
import Dropdown from '../Dropdown/Dropdown';

import * as actions from '../../../store/actions/index';

class IssueModal extends Component {

    state = {
        controls: {
            title: "",
            description: "",
            type: "Bug",
            priority: "Low"
        }
    }

    inputChangedHandler = (event, inputId) => {
        const updatedControls = {
            ...this.state.controls,
            [inputId]: event.target.value
        };
        this.setState({
            controls: updatedControls
        });
    }

    dropdownChangedHandler = (event) => {
        console.log(event)
        const updatedControls = {
            ...this.state.controls,
            type: event
        };
        this.setState({
            controls: updatedControls
        });
    }

    submitIssueHandler = (event) => {
        event.preventDefault()
        console.log("Issue Submit Called");
        let newIssue = {
            Title: this.state.controls.title,
            Description: this.state.controls.description,
            Type: this.state.controls.type,
            Priority: this.state.controls.priority,
            ProjectID: this.props.project.ProjectID,
            Creator: this.props.email
        }
        this.props.onAddIssue(this.props.token, newIssue);
        this.props.handleClose();
    }

    render() {
        const issueTypes = ['Bug', 'Feature', 'Modification'];
        const issuePrios = ['Low', 'Medium', 'High'];
        let content = (
            <Wrapper>
                <Modal show={this.props.show} onHide={this.props.onHide}>
                    <Modal.Header closeButton>
                        <Modal.Title>New Issue</Modal.Title>
                    </Modal.Header>
                    <Form onSubmit={this.submitIssueHandler}>
                        <Modal.Body>
                                <Form.Group controlId="formIssueTitle">
                                    <Form.Label>Issue Title</Form.Label>
                                    <Input 
                                        inputType="text" 
                                        inputPlaceholder="Enter Title"
                                        changed={(event) => this.inputChangedHandler(event, "title")}/>
                                </Form.Group>

                                <Form.Group controlId="formIssueDescription">
                                    <Form.Label>Description</Form.Label>
                                    <Input 
                                        inputType="text"
                                        inputAs="textarea"
                                        rows="3" 
                                        inputPlaceholder="Enter Description"
                                        changed={(event) => this.inputChangedHandler(event, "description")}/>
                                </Form.Group>

                                <Form.Group controlId="formIssueType">
                                    <Form.Label>Type</Form.Label>
                                    <Dropdown
                                        title={this.state.controls.type}
                                        items={issueTypes}
                                        onSelect={this.dropdownChangedHandler}/>
                                </Form.Group>

                                <Form.Group controlId="formIssueType">
                                    <Form.Label>Priority</Form.Label>
                                    <Dropdown
                                        title={this.state.controls.priority}
                                        items={issuePrios}
                                        onSelect={this.dropdownChangedHandler}/>
                                </Form.Group>
                        </Modal.Body>
                        <Modal.Footer>
                        <Button btnType="secondary" clicked={this.props.handleClose}>
                            Close
                        </Button>
                        <Button btnType="primary" type="submit">
                            Submit
                        </Button>
                        </Modal.Footer>
                    </Form>
                </Modal>
            </Wrapper>
        );

        return(
            <Wrapper>
                {content}
            </Wrapper>
        )
    }
}

const mapStateToProps = state => {
    return {
        token: state.auth.token,
        email: state.auth.email,
        project: state.projects.project
    }
}

const mapDispatchToProps = dispatch => {
    return {
        onAddIssue: (token, issueData) => dispatch(actions.addIssue(token, issueData))
    }
}

export default connect(mapStateToProps, mapDispatchToProps) (IssueModal);