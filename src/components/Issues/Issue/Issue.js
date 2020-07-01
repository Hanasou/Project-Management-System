import React, { Component } from 'react';
import { connect } from 'react-redux';

import classes from './Issue.module.css';
import Wrapper from '../../../hoc/Wrapper';
import Dropdown from '../../UI/Dropdown/Dropdown';

class Issue extends Component {
    render() {
        const statusItems = ['To Do', 'In Progress', 'Finished']
        return(
            <Wrapper>
                <div className={classes.Issue}>
                    <h5>{this.props.issue.Type}</h5>
                    <h1>{this.props.issue.Title}</h1>
                    <p>{this.props.issue.Description}</p>
                </div>
                <div style={{
                    float:'right',
                    marginRight:'20px'
                    }}>
                    <Dropdown 
                        title={this.props.issue.Status}
                        items={statusItems}/>
                </div>
            </Wrapper>
        )
    }
}

const mapStateToProps = state => {
    return {
        token: state.auth.token,
        issue: state.issues.issue
    };
};

export default connect(mapStateToProps) (Issue);