import React, {Component} from 'react';
import { connect } from 'react-redux';

import Wrapper from '../../../hoc/Wrapper';
import CustomModal from '../../../components/UI/Modal/CustomModal';
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

    
    render() {
        return (
            <Wrapper>
                <h1>
                    {this.props.project.Title}
                </h1>
                <Issues 
                    issues={this.props.issues}/> 
                <Issue />
                <CustomModal
                    type="addTeamMembers" 
                    show={this.state.show}
                    projectID={this.props.project.ProjectID}
                    projectName={this.props.project.Title}
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
        email: state.auth.email,
        issues: state.issues.issues,
        issue: state.issues.issue
    };
};

const mapDispatchToProps = dispatch => {
    return {
        onGetProject: (token, projectID, email) => dispatch(actions.getProject(token, projectID, email)),
        onGetIssues: (token, projectID) => dispatch(actions.getIssues(token, projectID)),
        onGetIssue: (token, projectID, issueID) => dispatch(actions.getIssue(token, projectID, issueID))
    }
}

export default connect( mapStateToProps, mapDispatchToProps) (Project);