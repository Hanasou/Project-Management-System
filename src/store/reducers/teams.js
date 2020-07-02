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
    return updateObject( state, {
        teams: [...state.teams, action.teamData]
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