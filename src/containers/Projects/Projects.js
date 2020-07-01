import React, { Component } from 'react';
import { connect } from 'react-redux';

import Card from '../../components/UI/Card/Card';
import Button from '../../components/UI/Button/Button';
import CustomModal from '../../components/UI/Modal/CustomModal';
import Wrapper from '../../hoc/Wrapper';
import classes from './Projects.module.css';
import * as actions from '../../store/actions/index';

class Projects extends Component {

    state = {
        show: false
    }
    
    componentDidMount() {
        this.props.onGetProjects(this.props.token, this.props.email);
    }

    handleShow = () => {
        this.setState({show: true});
    }

    handleClose = () => {
        this.setState({show: false});
    }

    render() {
        let projects = this.props.projects.map( project => (
            <Card
                key={project.ProjectID}
                cardType="Project" 
                title={project.Title}
                desc={project.Description}
                id={project.ProjectID}
                path={this.props.match.path}/>
        ))
        return(
            <Wrapper>
                <div className={classes.Projects}>
                    {projects}
                    <Button 
                        btnType='primary'
                        clicked={this.handleShow}
                        >New Project</Button>
                </div>
                <CustomModal
                    type="addProject" 
                    show={this.state.show}
                    onHide={this.handleClose}
                    handleClose={this.handleClose}/>
                
            </Wrapper>
        );
    };
};

const mapStateToProps = state => {
    return {
        projects: state.projects.projects,
        token: state.auth.token,
        email: state.auth.email
    }
}

const mapDispatchToProps = dispatch => {
    return {
        onGetProjects: (token, email) => dispatch(actions.getProjects(token, email))
    }
}

export default connect(mapStateToProps, mapDispatchToProps) (Projects);