import { combineReducers, legacy_createStore } from "redux";
import userReducer from "./user";


export const reducers = combineReducers({
    user: userReducer
})

export const globalStore = legacy_createStore(reducers);