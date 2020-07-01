import React, { Component } from 'react';
import { Route, Switch, withRouter, Redirect } from 'react-router-dom';
import { connect } from 'react-redux';
import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';

import Wrapper from './hoc/Wrapper';
import Signup from './containers/Auth/Signup';
import CustomNavbar from './containers/Navbar/Navbar';
import Projects from './containers/Projects/Projects';
import Project from './containers/Projects/Project/Project';
import Teams from './containers/Teams/Teams';
import * as actions from './store/actions/index';

class App extends Component {
  
  componentDidMount () {
    this.props.onTryAutoLogin();
  }

  render() {

    // Put routes in here later

    // Change content based on whether user is authenticated or not.
    let content = (
      <Wrapper>
        <CustomNavbar type="auth"/>
        <Signup/>
      </Wrapper>
    );

    if (this.props.isAuthenticated) {
      let routes = (
        <Switch>
          <Route exact path="/projects" component={Projects} />
          <Route exact path="/teams" component={Teams} />
          <Route path={"/projects/:projectID"} component={Project} />
          <Redirect to="/" />
        </Switch>
      )
      content = (
        <Wrapper>
          <CustomNavbar />
          { routes }
        </Wrapper>
      );
    }

    return (
      <Wrapper>
        {content}
      </Wrapper>
    );

  }
}

const mapStateToProps = state => {
  return {
    isAuthenticated: state.auth.token !== null
  };
};

const mapDispatchToProps = dispatch => {
  return {
    onTryAutoLogin: () => dispatch( actions.authCheckState() )
  };
};

export default withRouter( connect(mapStateToProps, mapDispatchToProps) (App) );
