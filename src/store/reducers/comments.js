import * as actionTypes from '../actions/actionTypes';
import { updateObject } from '../Utility';

const initialState = {
    comments: [],
}

const getCommentsSuccess = (state, action) => {
    return updateObject(state, {
        comments: action.comments
    });
}

const addCommentSuccess = (state, action) => {
    return updateObject(state, {
        comments: [...state.comments, action.commentData]
    })
}

const reducer = (state = initialState, action) => {
    switch (action.type) {
        case actionTypes.GET_ISSUES_SUCCESS: return getCommentsSuccess(state, action);
        case actionTypes.ADD_COMMENT_SUCCESS: return addCommentSuccess(state, action);
        default:
            return state;
    }
}

export default reducer;