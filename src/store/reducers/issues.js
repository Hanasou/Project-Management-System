import * as actionTypes from '../actions/actionTypes';
import { updateObject } from '../Utility';

const initialState = {
    issues: [],
    issue: {
        IssueID: "",
        Title: "",
        Description: "",
        Type: "",
        Status: "",
        Priority: ""
    }
}

const getIssuesSuccess = (state, action) => {
    return updateObject(state, {
        issues: action.issues
    });
}

const getIssueSuccess = (state, action) => {
    return updateObject(state, {
        issue: action.issue
    })
}

const addIssueSuccess = (state, action) => {
    return updateObject(state, {
        issues: [...state.issues, action.issueData]
    })
}

const updateIssueSuccess = (state, action) => {
    // Replace the issues array with the new issue
    let oldIssueKey = "";
    const newIssueKey = action.issueData.IssueID;
    const issueArray = [...state.issues];

    for (let i in issueArray) {
        oldIssueKey = issueArray[i].IssueID;
        if (oldIssueKey === newIssueKey) {
            issueArray[i] = action.issueData;
            break;
        }
    }

    console.log("Updated Issues ", issueArray)
    return updateObject(state, {
        issue: action.issueData,
        issues: issueArray
    })
}

const deleteIssueSuccess = (state, action) => {
    const defaultIssue = {
        IssueID: "",
        Title: "",
        Description: "",
        Type: "",
        Status: "",
        Priority: ""
    }
    const filterIssue = state.issues.filter(issue => issue.IssueID !== action.deleteRequest.IssueID)
    return updateObject(state, {
        issue: defaultIssue,
        issues: filterIssue
    })
}

const reducer = (state = initialState, action) => {
    switch (action.type) {
        case actionTypes.GET_ISSUES_SUCCESS: return getIssuesSuccess(state, action);
        case actionTypes.GET_ISSUE_SUCCESS: return getIssueSuccess(state, action);
        case actionTypes.ADD_ISSUE_SUCCESS: return addIssueSuccess(state, action);
        case actionTypes.UPDATE_ISSUE_SUCCESS: return updateIssueSuccess(state, action);
        case actionTypes.DELETE_ISSUE_SUCCESS: return deleteIssueSuccess(state, action);
        default:
            return state;
    }
}

export default reducer;