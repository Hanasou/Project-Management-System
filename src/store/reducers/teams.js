import * as actionTypes from '../actions/actionTypes';
import {updateObject} from '../Utility';

const initialState = {
    teams: []
}

const getTeamsSuccess = ( state, action ) => {
    return updateObject( state, {
        teams: action.teams
    });
}

const addTeamSuccess = (state, action) => {
    let currTeams = [...state.teams]
    const projectID = action.teamData.ProjectID
    for (let i in currTeams) {
        const currTeam = currTeams[i]
        if (currTeam.ProjectID === projectID) {
            currTeam.Members.push(action.teamData.Members[0])
        }
    }
    return updateObject( state, {
        teams: currTeams
    })
}

const reducer = ( state = initialState, action ) => {
    switch ( action.type ) {
        case actionTypes.GET_TEAMS_SUCCESS: return getTeamsSuccess(state, action);
        case actionTypes.ADD_TEAM_SUCCESS: return addTeamSuccess(state, action);
        default:
            return state;
    }
};

export default reducer;