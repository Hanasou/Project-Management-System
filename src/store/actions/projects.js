import * as actionTypes from './actionTypes';
import axios from '../../axios';

export const getProjectsSuccess = ( projects ) => {
    return {
        type: actionTypes.GET_PROJECTS_SUCCESS,
        projects: projects
    }
}

export const getProjectsFail = ( error ) => {
    return {
        type: actionTypes.GET_PROJECTS_FAIL,
        error: error
    }
}

export const getProjects = ( token, email ) => {
    return dispatch => {
        const headers = {
            'Authorization': token
        }
        axios.get('/projects/getAll/' + email, {headers})
            .then(response => {
                console.log(response)
                const fetchedProjects = [];
                for (let key in response.data) {
                    fetchedProjects.push({
                        ...response.data[key],
                        Id: key
                    });
                }
                dispatch(getProjectsSuccess(fetchedProjects));
            })
            .catch(error => {
                dispatch(getProjectsFail(error));
            })
    }
}

export const addProjectSuccess = ( projectData ) => {
    return {
        type: actionTypes.ADD_PROJECT_SUCCESS,
        projectData: projectData
    }
}

export const addProjectFail = ( error ) => {
    return {
        type: actionTypes.ADD_PROJECT_FAIL,
        error: error
    }
}

export const addProject = ( token, email, title, description) => {
    return dispatch => {
        const headers = {
            'Authorization': token
        };
        const projectData = {
            Title: title,
            Description: description
        };
        console.log(projectData)
        axios.post('/projects/add/' + email, projectData, {headers})
            .then(response => {
                console.log(response.data);
                dispatch(addProjectSuccess(projectData));
            })
            .catch( error => {
                dispatch(addProjectFail(error))
            })
    }
}