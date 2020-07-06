import * as actionTypes from './actionTypes';
import axios from '../../axios';

export const getIssuesSuccess = (issues) => {
    return {
        type: actionTypes.GET_ISSUES_SUCCESS,
        issues: issues
    }
}

export const getIssuesFail = (error) => {
    return {
        type: actionTypes.GET_ISSUES_FAIL,
        error: error
    }
}

export const getIssues = (token, projectID) => {
    return dispatch => {
        const headers = {
            'Authorization': token
        };
        axios.get('/issues/getAll/'+projectID, {headers})
            .then(response => {
                console.log(response.data);
                const fetchedIssues = [];
                for (let key in response.data) {
                    fetchedIssues.push({
                        ...response.data[key],
                        Id: key
                    });
                }
                dispatch(getIssuesSuccess(fetchedIssues));
            })
            .catch(error => {
                dispatch(getIssuesFail(error));
            })
    }
}

export const getIssueSuccess = (issue) => {
    return {
        type: actionTypes.GET_ISSUE_SUCCESS,
        issue: issue
    }
}

export const getIssueFail = (error) => {
    return {
        type: actionTypes.GET_ISSUE_FAIL,
        error: error
    }
}

export const getIssue = (token, projectID, issueID) => {
    return dispatch => {
        const headers = {
            'Authorization': token
        }
        axios.get('/issues/get/'+projectID+'/'+issueID, {headers})
            .then(response => {
                const fetchedIssue = response.data;
                console.log(fetchedIssue);
                dispatch(getIssueSuccess(fetchedIssue));
            })
            .catch(error => {
                dispatch(getIssueFail(error));
            })
    }
}

export const addIssueSuccess = ( issueData ) => {
    return {
        type: actionTypes.ADD_ISSUE_SUCCESS,
        issueData: issueData
    }
}

export const addIssueFail = ( error ) => {
    return {
        type: actionTypes.ADD_ISSUE_FAIL,
        error: error
    }
}

export const addIssue = (token, issueData) => {
    return dispatch => {
        const headers = {
            'Authorization': token
        }
        axios.post('/issues/add', issueData, {headers})
            .then(response => {
                console.log("Added Issue: ", response.data);
                dispatch(addIssueSuccess(response.data));
            })
            .catch(error => {
                console.log(error)
                dispatch(addIssueFail(error));
            })
    }
}