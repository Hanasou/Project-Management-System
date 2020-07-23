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
        console.log("Project ID: ", projectID);
        console.log("Issue ID: ", issueID);
        axios.get('/issues/get/'+projectID+'/'+issueID, {headers})
            .then(response => {
                const fetchedIssue = response.data;
                console.log(fetchedIssue);
                dispatch(getIssueSuccess(fetchedIssue));
            })
            .catch(error => {
                console.log(error);
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

export const updateIssueSuccess = (issueData) => {
    return {
        type: actionTypes.UPDATE_ISSUE_SUCCESS,
        issueData: issueData
    }
}

export const updateIssueFail = (error) => {
    return {
        type: actionTypes.UPDATE_ISSUE_FAIL,
        error: error
    }
}

export const updateIssue = (token, updateData) => {
    return dispatch => {
        const headers = {
            'Authorization': token
        }
        axios.post('/issues/update', updateData, {headers})
            .then(response => {
                console.log("Updated Data: ", response.data);
                dispatch(updateIssueSuccess(response.data));
            })
            .catch(error => {
                console.log(error);
                dispatch(updateIssueFail(error));
            })
    }
}

export const deleteIssueSuccess = (deleteRequest) => {
    return {
        type: actionTypes.DELETE_ISSUE_SUCCESS,
        deleteRequest: deleteRequest
    }
}

export const deleteIssueFail = (error) => {
    return {
        type: actionTypes.DELETE_ISSUE_FAIL,
        error: error
    }
}

export const deleteIssue = (token, deleteRequest) => {
    return dispatch => {
        const headers = {
            'Authorization': token
        }
        const issueID = deleteRequest.IssueID;
        const projectID = deleteRequest.ProjectID;
        axios.delete('/issues/delete/' + issueID + '/' + projectID, {headers})
            .then(response => {
                console.log("Delete Request: ", response.data);
                dispatch(deleteIssueSuccess(response.data));
            })
            .catch(error => {
                console.log(error);
                dispatch(deleteIssueFail(error));
            })
    }
}