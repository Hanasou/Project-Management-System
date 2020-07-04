import React, { Component } from 'react';
import { connect } from 'react-redux';
import Modal from 'react-bootstrap/Modal';
import Form from 'react-bootstrap/Form';

import Wrapper from '../../../hoc/Wrapper';
import Button from '../Button/Button';
import Input from '../Input/Input';

import * as actions from '../../../store/actions/index';

class CustomModal extends Component {

    state = {
        controls: {
            title: "",
            description: "",
            members: ""
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

    submitProjectHandler = (event) => {
        event.preventDefault()
        console.log("Project Submit Called");
        this.props.onAddProject(this.props.token, this.props.email, 
            this.state.controls.title, this.state.controls.description);
        this.props.handleClose();
    }

    submitTeamMembersHandler = (event) => {
        event.preventDefault()
        console.log("Submit team members called")
        
        const membersArray = this.state.controls.members.split(" ")
        const requestData = {
            ProjectID: this.props.projectID,
            Members: membersArray,
            ProjectName: this.props.projectName,
            ProjectDesc: this.props.projectDesc
        }
        this.props.onAddTeamMembers(this.props.token, requestData)
        this.props.handleClose()
    }

    render() {
        let content = (
            <Wrapper>
                <Modal show={this.props.show} onHide={this.props.onHide}>
                    <Modal.Header closeButton>
                    <Modal.Title>Modal heading</Modal.Title>
                    </Modal.Header>
                    <Modal.Body>Woohoo, you're reading this text in a modal!</Modal.Body>
                    <Modal.Footer>
                    <Button btnType="secondary" clicked={this.props.handleClose}>
                        Close
                    </Button>
                    <Button btnType="primary" clicked={this.props.handleClose}>
                        Save Changes
                    </Button>
                    </Modal.Footer>
                </Modal>
            </Wrapper>
        );
        if (this.props.type === "addProject") {
            content = (
                <Wrapper>
                    <Modal show={this.props.show} onHide={this.props.onHide}>
                        <Modal.Header closeButton>
                            <Modal.Title>New Project</Modal.Title>
                        </Modal.Header>
                        <Form onSubmit={this.submitProjectHandler}>
                            <Modal.Body>
                                    <Form.Group controlId="formProjectTitle">
                                        <Form.Label>Project Title</Form.Label>
                                        <Input 
                                            inputType="text" 
                                            inputPlaceholder="Enter Title"
                                            changed={(event) => this.inputChangedHandler(event, "title")}/>
                                    </Form.Group>

                                    <Form.Group controlId="formProjectDescription">
                                        <Form.Label>Description</Form.Label>
                                        <Input 
                                            inputType="text"
                                            inputAs="textarea"
                                            rows="3" 
                                            inputPlaceholder="Enter Description"
                                            changed={(event) => this.inputChangedHandler(event, "description")}/>
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
            )
        }
        else if (this.props.type === 'addTeamMembers') {
            content = (
                <Wrapper>
                    <Modal show={this.props.show} onHide={this.props.onHide}>
                        <Modal.Header closeButton>
                            <Modal.Title>{this.props.projectName}</Modal.Title>
                        </Modal.Header>
                        <Form onSubmit={this.submitTeamMembersHandler}>
                            <Modal.Body>
                                <Form.Group controlId="formMembers">
                                    <Form.Label>Add Members</Form.Label>
                                    <Input 
                                        inputType="text" 
                                        inputPlaceholder="Add Members"
                                        changed={(event) => this.inputChangedHandler(event, "members")}/>
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
            )
        }

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
        email: state.auth.email
    }
}

const mapDispatchToProps = dispatch => {
    return {
        onAddProject: (token, email, title, description) => 
            dispatch(actions.addProject(token, email, title, description)),
        onAddTeamMembers: (token, teamData) => dispatch(actions.addTeam(token, teamData))
    }
}

export default connect(mapStateToProps, mapDispatchToProps) (CustomModal);