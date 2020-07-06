import * as actionTypes from './actionTypes';
import axios from '../../axios';

export const getCommentsSuccess = (comments) => {
    return {
        type: actionTypes.GET_COMMENTS_SUCCESS,
        comments: comments
    }
}

export const getCommentsFail = (error) => {
    return {
        type: actionTypes.GET_COMMENTS_FAIL,
        error: error
    }
}

export const getComments = (token, issueID) => {
    return dispatch => {
        const headers = {
            'Authorization': token
        };
        axios.get('/comments/getAll/'+issueID, {headers})
            .then(response => {
                console.log(response.data);
                const fetchedComments = [];
                for (let key in response.data) {
                    fetchedComments.push({
                        ...response.data[key],
                        Id: key
                    });
                }
                dispatch(getCommentsSuccess(fetchedComments));
            })
            .catch(error => {
                dispatch(getCommentsFail(error));
            })
    }
}

export const addCommentSuccess = ( commentData ) => {
    return {
        type: actionTypes.ADD_COMMENT_SUCCESS,
        issueData: commentData
    }
}

export const addCommentFail = ( error ) => {
    return {
        type: actionTypes.ADD_COMMENT_FAIL,
        error: error
    }
}

export const addComment = (token, commentData) => {
    return dispatch => {
        const headers = {
            'Authorization': token
        }
        axios.post('/issues/add', commentData, {headers})
            .then(response => {
                console.log("Added Comment: ", response.data);
                dispatch(addCommentSuccess(response.data));
            })
            .catch(error => {
                dispatch(addCommentFail(error));
            })
    }
}