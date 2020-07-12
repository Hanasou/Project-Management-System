import React, { Component } from 'react';
import { connect } from 'react-redux';
import ListGroup from 'react-bootstrap/ListGroup';

import classes from './Issues.module.css';
import Button from '../UI/Button/Button';
import Wrapper from '../../hoc/Wrapper';
import IssueModal from '../UI/Modal/IssueModal';
import * as actions from '../../store/actions/index';

class Issues extends Component {

    state = {
        show: false
    }
    
    onIssueClicked = (event, issueID) => {
        event.preventDefault();
        this.props.onGetIssue(this.props.token,this.props.project.ProjectID,issueID);
        this.props.onGetComments(this.props.token, issueID);
    }

    handleShow = () => {
        this.setState({show: true});
    }

    handleClose = () => {
        this.setState({show: false});
    }

    render () {
        let issues = this.props.issues.map( issue => (
            <ListGroup.Item 
                key={issue.IssueID}
                action
                onClick={(event) => this.onIssueClicked(event, issue.IssueID)}>
                {issue.Title}
            </ListGroup.Item>
        ))

        return(
            <Wrapper>
                <div className={classes.Issues}> 
                    <h2>Issues</h2>
                    <ListGroup className={classes.IssuesList}>
                        {issues}
                    </ListGroup>
                    <div style={{
                        marginTop: '20px'
                        }
                    }>
                        <Button 
                            btnType="success"
                            clicked={this.handleShow}
                            >Add Issue</Button>
                    </div>
                </div>
                <IssueModal
                    show={this.state.show}
                    onHide={this.handleClose}
                    handleClose={this.handleClose}/>
            </Wrapper>
        )
    }
}

const mapStateToProps = state => {
    return {
        token: state.auth.token,
        project: state.projects.project,
        issue: state.issues.issue,
        comments: state.comments.comments
    };
};

const mapDispatchToProps = dispatch => {
    return {
        onGetIssue: (token, projectID, issueID) => dispatch(actions.getIssue(token, projectID, issueID)),
        onGetComments: (token, issueID) => dispatch(actions.getComments(token, issueID))
    }
}

export default connect(mapStateToProps, mapDispatchToProps) (Issues);