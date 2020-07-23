import React, {Component} from 'react';
import { connect } from 'react-redux';

import Wrapper from '../../../hoc/Wrapper';
import Button from '../../../components/UI/Button/Button';
import Issues from '../../../components/Issues/Issues';
import Issue from '../../../components/Issues/Issue/Issue';
import * as actions from '../../../store/actions/index';

class Project extends Component {

    state = {
        show: false
    }

    componentDidMount() {
        const projectID = this.props.match.params.projectID;
        const userEmail = this.props.email;
        this.props.onGetProject(this.props.token, projectID, userEmail);
        this.props.onGetIssues(this.props.token, projectID);
    }

    handleShow = () => {
        this.setState({show: true})
    }

    handleClose = () => {
        this.setState({show: false})
    }

    handleDelete = (event) => {
        event.preventDefault();
        const deleteRequest = {
            ProjectID: this.props.project.ProjectID,
            Email: this.props.email
        };
        this.props.onDeleteProject(this.props.token, deleteRequest);
    }

    
    render() {
        let issue = null;
        if (this.props.issue.IssueID !== "") {
            issue = <Issue/>
        }
        return (
            <Wrapper>
                <div>
                    <h1>
                        {this.props.project.Title}
                    </h1>
                    <Button
                        btnType="danger"
                        clicked={this.handleDelete}>
                        Leave Project
                    </Button>
                </div>
                <Issues 
                    issues={this.props.issues}/> 
                {issue}
            </Wrapper>
        )
    }
}

const mapStateToProps = state => {
    return {
        token: state.auth.token,
        project: state.projects.project,
        email: state.auth.email,
        issues: state.issues.issues,
        issue: state.issues.issue
    };
};

const mapDispatchToProps = dispatch => {
    return {
        onGetProject: (token, projectID, email) => dispatch(actions.getProject(token, projectID, email)),
        onGetIssues: (token, projectID) => dispatch(actions.getIssues(token, projectID)),
        onGetIssue: (token, projectID, issueID) => dispatch(actions.getIssue(token, projectID, issueID)),
        onDeleteProject: (token, deleteRequest) => dispatch(actions.deleteProject(token, deleteRequest))
    }
}

export default connect( mapStateToProps, mapDispatchToProps) (Project);