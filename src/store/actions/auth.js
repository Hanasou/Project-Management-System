import * as actionTypes from './actionTypes';
import axios from '../../axios';

export const authStart = () => {
    return {
        type: actionTypes.AUTH_START
    };
};

export const authSuccess = (token, email) => {
    return {
        type: actionTypes.AUTH_SUCCESS,
        token: token,
        email: email
    };
};

export const authFail = (error) => {
    return {
        type: actionTypes.AUTH_FAIL,
        error: error
    };
};

export const checkAuthTimeout = (expirationTime) => {
    return dispatch => {
        setTimeout(() => {
            dispatch(authLogout());
        }, expirationTime * 1000);
    };
};

export const authSignup = (email, password) => {
    return dispatch => {
        dispatch(authStart())
        const signupData = {
            Email: email,
            Password: password
        }
        console.log(signupData)
        axios.post('/auth/signup', signupData)
            .then(response => {
                console.log(response)
            })
            .catch(err => {
                dispatch(authFail(err.response.data.error))
            })
    }
}

export const authLogin = (email, password) => {
    return dispatch => {
        dispatch(authStart())
        const loginData = {
            Email: email,
            Password: password
        }
        axios.post('/auth/login', loginData)
            .then(response => {
                console.log(response);
                const expirationDate = new Date(new Date().getTime() + response.data.expiresIn);
                console.log(expirationDate)
                localStorage.setItem('token', response.data.token);
                localStorage.setItem('expirationDate', expirationDate);
                localStorage.setItem('email', response.data.userEmail);
                dispatch(authSuccess(response.data.token, response.data.userEmail));
                //dispatch(checkAuthTimeout(response.data.expiresIn));
            })
            .catch(err => {
                dispatch(authFail(err.response.data.error))
            })
    }
}

export const authLogout = () => {
    localStorage.removeItem("token");
    localStorage.removeItem("expirationDate");
    localStorage.removeItem("email");
    return {
        type: actionTypes.AUTH_LOGOUT
    };
};

export const setAuthRedirectPath = (path) => {
    return {
        type: actionTypes.SET_AUTH_REDIRECT_PATH,
        path: path
    };
};

export const authCheckState = () => {
    return dispatch => {
        const token = localStorage.getItem('token');
        if (!token) {
            dispatch(authLogout());
        } else {
            const expirationDate = new Date(localStorage.getItem('expirationDate'));
            if (expirationDate <= new Date()) {
                dispatch(authLogout());
            } else {
                const email = localStorage.getItem('email');
                dispatch(authSuccess(token, email));
                dispatch(checkAuthTimeout((expirationDate.getTime() - new Date().getTime()) / 1000 ));
            }   
        }
    };
};