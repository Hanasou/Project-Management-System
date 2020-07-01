import * as actionTypes from '../actions/actionTypes';
import { updateObject } from '../Utility';

const initialState = {
    projects: [],
    project: {
        Title: "",
        Email: "",
        Description: "",
        ProjectID: ""
    }
}

const getProjectsSuccess = ( state, action ) => {
    return updateObject( state, {
        projects: action.projects
    });
}

const getProjectSuccess = (state, action) => {
    return updateObject( state, {
        project: action.project
    })
}

const addProjectSuccess = (state, action) => {
    return updateObject( state, {
        projects: [...state.projects, action.projectData]
    })
}

const reducer = ( state = initialState, action ) => {
    switch ( action.type ) {
        case actionTypes.GET_PROJECTS_SUCCESS: return getProjectsSuccess(state, action);
        case actionTypes.GET_PROJECT_SUCCESS: return getProjectSuccess(state, action);
        case actionTypes.ADD_PROJECT_SUCCESS: return addProjectSuccess(state, action);
        default:
            return state;
    }
};

export default reducer;