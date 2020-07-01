import React, {Component} from 'react';
import { connect } from 'react-redux';

import Wrapper from '../../hoc/Wrapper';

class Teams extends Component {
    render() {
        return(
            <Wrapper>
                <h1>This is the teams page</h1>
            </Wrapper>
        )
    }
}

const mapStateToProps = state => {
    return {
        token: state.auth.token,
        projects: state.projects.projects
    };
};

export default connect(mapStateToProps)(Teams);