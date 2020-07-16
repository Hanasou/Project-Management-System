import React, { Component } from 'react';
import { connect } from 'react-redux';

import classes from './Issue.module.css';
import Wrapper from '../../../hoc/Wrapper';
import Input from '../../UI/Input/Input';
import Button from '../../UI/Button/Button';
import Comments from '../../Comments/Comments';
import Dropdown from '../../UI/Dropdown/Dropdown';
import * as actions from '../../../store/actions/index';

class Issue extends Component {

    state = {
        titleEdit: false,
        descEdit: false,
        changed: false,
        dropdowns: {
            statusTitle: this.props.issue.Status,
            prioTitle: this.props.issue.Priority
        },
        tags: {
            title: this.props.issue.Title,
            description: this.props.issue.Description
        }
    }

    componentDidMount() {
        document.addEventListener("mousedown", this.handleClick, false)
    }

    componentWillUnmount() {
        document.removeEventListener("mousedown", this.handleClick, false)
    }

    onChangeHandler = (event, inputId) => {
        const changedTags = {
            ...this.state.tags,
            [inputId]: event.target.value
        };
        this.setState({
            ...this.state.tags,
            tags: changedTags
        })
    }

    tagClickHandler = (event, tag) => {
        event.preventDefault()
        this.setState({
            [tag]: true,
            changed: true
        })
    }

    onDropdownItemClick = (event, inputId) => {
        console.log(event)
        const updatedDropdowns = {
            ...this.state.dropdowns,
            [inputId]: event
        }
        this.setState({
            ...this.state.dropdowns,
            changed: true,
            dropdowns: updatedDropdowns
        })
    }

    handleClick = (event) => {
        if (this.node.contains(event.target) || this.node2.contains(event.target)) {
            return;
        }
        const originalDropdowns = {
            statusTitle: this.props.issue.Status,
            prioTitle: this.props.issue.Priority
        };
        const originalTags = {
            title: this.props.issue.Title,
            description: this.props.issue.Description
        }
        this.setState({
            ...this.state,
            typeEdit: false,
            titleEdit: false,
            descEdit: false,
            changed: false,
            dropdowns: originalDropdowns,
            tags: originalTags
        })
    }

    handleSave = (event) => {
        event.preventDefault();
        const updateIssueData = {
            IssueID: this.props.issue.IssueID,
            ProjectID: this.props.issue.ProjectID,
            Title: this.state.tags.title,
            Description: this.state.tags.description,
            Status: this.state.dropdowns.statusTitle,
            Priority: this.state.dropdowns.prioTitle
        };
        this.handleClick(event)
        const token = this.props.token;
        this.props.onUpdateIssue(token, updateIssueData);
    }

    render() {
        const statusItems = ['To Do', 'In Progress', 'Finished']
        const prioItems = ['Low', 'Medium', 'High']

        let titleTag = null;
        let descTag = null;
        let changeButton = null;

        if (this.state.changed === true) {
            changeButton = (
                <Button
                    btnType="success"
                    clicked={this.handleSave}>
                        Save Changes
                </Button>
            )
        }

        if (this.state.titleEdit === false) {
            titleTag = <h5
                        onClick={(event) => this.tagClickHandler(event, "titleEdit")}>
                            {this.props.issue.Title}
                        </h5>
        } else {
            titleTag = (
                <Input
                    changed={(event) => this.onChangeHandler(event, "title")}
                    inputType="text">
                        {this.props.issue.Title}
                </Input>
            )
        }

        if (this.state.descEdit === false) {
            descTag = <p
                        onClick={(event) => this.tagClickHandler(event, "descEdit")}>
                            {this.props.issue.Description}
                        </p>
        } else {
            descTag = (
                <Input
                    changed={(event) => this.onChangeHandler(event, "description")}
                    inputType="text"
                    inputAs="textarea"
                    rows="3">
                        {this.props.issue.Description}
                </Input>
            )
        }

        return(
            <Wrapper>
                <div 
                    className={classes.Issue}
                    ref={node => this.node = node}>
                    <h5>{this.props.issue.Type}</h5>
                    {titleTag}
                    <br/>
                    {descTag}
                    <br/>
                    {changeButton}
                    <Comments/>
                </div>
                <div style={{
                    float:'right',
                    marginRight:'20px'
                    }}
                    ref={node2 => this.node2 = node2}>
                    <Dropdown 
                        title={this.state.dropdowns.statusTitle}
                        items={statusItems}
                        onItemClick={(event) => this.onDropdownItemClick(event, "statusTitle")}/>
                    <br/>
                    <Dropdown 
                        title={this.state.dropdowns.prioTitle}
                        items={prioItems}
                        onItemClick={(event) => this.onDropdownItemClick(event, "prioTitle")}/>
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

const mapDispatchToProps = dispatch => {
    return {
        onUpdateIssue: (token, updateData) => dispatch(actions.updateIssue(token, updateData))
    }
}

export default connect(mapStateToProps, mapDispatchToProps) (Issue);