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
                console.log(response.data)
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

export const getProjectSuccess = ( project ) => {
    return {
        type: actionTypes.GET_PROJECT_SUCCESS,
        project: project
    }
}

export const getProject = ( token, projectID, userEmail) => {
    return dispatch => {
        const headers = {
            'Authorization': token
        }
        axios.get('/projects/get/' + projectID + '/' + userEmail, {headers})
            .then(response => {
                console.log(response.data);
                const project = response.data;
                dispatch(getProjectSuccess(project));
            })
            .catch()
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
        axios.post('/projects/add/' + email, projectData, {headers})
            .then(response => {
                console.log("Added Project:", response.data);
                dispatch(addProjectSuccess(response.data));
            })
            .catch( error => {
                dispatch(addProjectFail(error))
            })
    }
}

export const deleteProjectSuccess = ( deleteRequest ) => {
    return {
        type: actionTypes.DELETE_PROJECT_SUCCESS,
        deleteRequest: deleteRequest
    }
}

export const deleteProjectFail = ( error ) => {
    return {
        type: actionTypes.DELETE_ISSUE_FAIL,
        error: error
    }
}

export const deleteProject = (token, deleteRequest) => {
    return dispatch => {
        const headers = {
            'Authorization': token
        };
        console.log(deleteRequest);
        const projectID = deleteRequest.ProjectID;
        const email = deleteRequest.Email;
        axios.delete('/projects/delete/' + projectID + '/' + email, {headers})
            .then(response => {
                console.log("Delete Request: ", response.data);
                dispatch(deleteProjectSuccess(response.data));
            })
            .catch(error => {
                console.log("Delete Request Error: ", error);
                dispatch(deleteProjectFail(error));
            })
    }
}