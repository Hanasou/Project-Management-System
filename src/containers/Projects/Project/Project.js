import React, {Component} from 'react';
import { connect } from 'react-redux';

import Wrapper from '../../../hoc/Wrapper';
import Issues from '../../../components/Issues/Issues';
import Issue from '../../../components/Issues/Issue/Issue';
import * as actions from '../../../store/actions/index';

class Project extends Component {


    componentDidMount() {
        const projectID = this.props.match.params.projectID;
        this.props.onGetProject(this.props.token, projectID);
        this.props.onGetIssues(this.props.token, projectID);
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
            </Wrapper>
        )
    }
}

const mapStateToProps = state => {
    return {
        token: state.auth.token,
        project: state.projects.project,
        issues: state.issues.issues,
        issue: state.issues.issue
    };
};

const mapDispatchToProps = dispatch => {
    return {
        onGetProject: (token, projectID) => dispatch(actions.getProject(token, projectID)),
        onGetIssues: (token, projectID) => dispatch(actions.getIssues(token, projectID)),
        onGetIssue: (token, projectID, issueID) => dispatch(actions.getIssue(token, projectID, issueID))
    }
}

export default connect( mapStateToProps, mapDispatchToProps) (Project);