import React, {Component} from 'react';
import { connect } from 'react-redux';

import Wrapper from '../../hoc/Wrapper';
import Card from '../../components/UI/Card/Card';
import CustomModal from '../../components/UI/Modal/CustomModal';
import * as actions from '../../store/actions/index';

class Teams extends Component {

    state = {
        show: false,
        selectedProject: "",
        selectedProjectName: ""
    }

    componentDidMount() {
        this.props.onGetTeams(this.props.token, this.props.email)
    }

    handleShow = (projectID, projectName) => {
        this.setState({
            show: true, 
            selectedProject: projectID,
            selectedProjectName: projectName
        });
    }

    handleClose = () => {
        this.setState({
            show: false, 
            selectedProject: "",
            selectedProjectName: ""
        });
    }

    render() {
        let teams = this.props.teams.map(team => (
            <Card
                key={team.ProjectID}
                cardType='Team'
                title={team.Project}
                id={team.ProjectID}
                path={this.props.match.path}
                members={team.Members}
                clicked={(projectID, projectName) => 
                    this.handleShow(team.ProjectID, team.Project)}/>
        ))
        return(
            <Wrapper>
                <h3>My Teams</h3>
                {teams}
                <CustomModal
                    type="addTeamMembers" 
                    show={this.state.show}
                    projectID={this.state.selectedProject}
                    projectName={this.state.selectedProjectName}
                    onHide={this.handleClose}
                    handleClose={this.handleClose}/>
            </Wrapper>
        )
    }
}

const mapStateToProps = state => {
    return {
        token: state.auth.token,
        email: state.auth.email,
        teams: state.teams.teams
    };
};

const mapDispatchToProps = dispatch => {
    return {
        onGetTeams: (token, email) => dispatch(actions.getTeams(token, email))
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(Teams);