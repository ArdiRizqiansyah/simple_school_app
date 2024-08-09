import Cookies from 'js-cookie';

const DEFAULT_STATE = {
    "name": "",
    "email": "",
    "token": "",
}

export default function userReducer(state = DEFAULT_STATE, action) {
    switch (action.type) {
        case "LOGIN":
            return {
                ...state,
                "name": action.payload.name,
                "email": action.payload.email,
                "token": action.payload.token,
            };
        case "LOGOUT":
            // hapus cookie
            Cookies.remove('user');
            Cookies.remove('token');

            return DEFAULT_STATE;
        default:
            return state;
    }
}