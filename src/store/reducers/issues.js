import * as actionTypes from '../actions/actionTypes';
import { updateObject } from '../Utility';

const initialState = {
    issues: [],
    issue: {
        IssueID: "",
        Title: "",
        Description: "",
        Type: "",
        Status: ""
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

const reducer = (state = initialState, action) => {
    switch (action.type) {
        case actionTypes.GET_ISSUES_SUCCESS: return getIssuesSuccess(state, action);
        case actionTypes.GET_ISSUE_SUCCESS: return getIssueSuccess(state, action);
        case actionTypes.ADD_ISSUE_SUCCESS: return addIssueSuccess(state, action);
        default:
            return state;
    }
}

export default reducer;