import * as actionTypes from './actionTypes';
import axios from '../../axios';

export const getTeamsSuccess = ( teams ) => {
    return {
        type: actionTypes.GET_TEAMS_SUCCESS,
        teams: teams
    }
}

export const getTeamsFail = ( error ) => {
    return {
        type: actionTypes.GET_TEAMS_FAIL,
        error: error
    }
}

export const getTeams = (token, email) => {
    return dispatch => {
        const headers = {
            'Authorization': token
        }
        axios.get('/teams/getAll/' + email, {headers})
            .then(response => {
                console.log(response.data)
                const fetchedTeams = []
                for (let key in response.data) {
                    fetchedTeams.push({
                        ...response.data[key],
                        Id: key
                    })
                }
                dispatch(getTeamsSuccess(fetchedTeams))
            })
            .catch(error => {
                dispatch(getTeamsFail(error))
            })
    }
}

export const addTeamSuccess = ( teamData ) => {
    return {
        type: actionTypes.ADD_TEAM_SUCCESS,
        teamData: teamData
    }
}

export const addTeamFail = ( error ) => {
    return {
        type: actionTypes.ADD_TEAM_FAIL,
        error: error
    }
}

export const addTeam = (token, teamData) => {
    return dispatch => {
        const headers = {
            'Authorization': token
        }
        axios.post('/teams/add', teamData, {headers})
            .then(response => {
                console.log(response.data)
                dispatch(addTeamSuccess(response.data))
            })
            .catch(error => {
                dispatch(addTeamFail(error))
            })
    }
}