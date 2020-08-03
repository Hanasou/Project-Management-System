import React, {Component} from 'react';
import { connect } from 'react-redux';
import Table from 'react-bootstrap/Table';

import Wrapper from '../../hoc/Wrapper';
import * as actions from '../../store/actions/index';

class HomeIssues extends Component {

    componentDidMount() {
        this.props.onGetIssues(this.props.token, this.props.email);
    }

    render() {
        const header = (
            <thead>
                <tr>
                    <th>Title</th>
                    <th>Type</th>
                    <th>Priority</th>
                    <th>Status</th>
                </tr>
            </thead>
        )
        let issueEntries = this.props.issues.map(issue => (
            <tr>
                <td>{issue.Title}</td>
                <td>{issue.Type}</td>
                <td>{issue.Priority}</td>
                <td>{issue.Status}</td>
            </tr>
        ));

        const body = (
            <tbody>
                {issueEntries}
            </tbody>
        )

        return(
            <Wrapper>
                <Table striped bordered hover>
                    {header}
                    {body}
                </Table>
            </Wrapper>
        )
    }
}

const mapStateToProps = state => {
    return {
        token: state.auth.token,
        email: state.auth.email,
        issues: state.issues.issues
    };
};

const mapDispatchToProps = dispatch => {
    return {
        onGetIssues: (token, userEmail) => dispatch(actions.getIssuesByUser(token, userEmail))
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(HomeIssues);