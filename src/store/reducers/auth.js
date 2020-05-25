import * as actionTypes from '../actions/actionTypes';
import { updateObject } from '../Utility';

const initialState = {
    token: null,
    email: null,
    error: null,
    authRedirectPath: "/"
}

const authStart = ( state, action ) => {
    return updateObject( state, { error: null });
};

const authSuccess = ( state, action ) => {
    return updateObject( state, {
        token: action.token,
        email: action.email,
        error: null
    });
};

const authFail = ( state, action ) => {
    return updateObject( state, {
        error: action.error
    });
};

const authLogout = (state, action) => {
    return updateObject(state, { token: null, email: null });
};

const setAuthRedirectPath = (state, action) => {
    return updateObject(state, { authRedirectPath: action.path });
};

const reducer = ( state = initialState, action ) => {
    switch ( action.type ) {
        case actionTypes.AUTH_START: return authStart(state, action);
        case actionTypes.AUTH_SUCCESS: return authSuccess(state, action);
        case actionTypes.AUTH_FAIL: return authFail(state, action);
        case actionTypes.AUTH_LOGOUT: return authLogout(state, action);
        case actionTypes.SET_AUTH_REDIRECT_PATH: return setAuthRedirectPath(state, action);
        default:
            return state;
    }
};

export default reducer;