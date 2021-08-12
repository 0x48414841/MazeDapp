import { combineReducers } from "redux";
import mazeReducer  from "./mazeReducer";
import player1LocationReducer from "./player1LocationReducer";
import usernameReducer from "./usernameReducer";

export default combineReducers({
    maze: mazeReducer,
    playersLoc: player1LocationReducer,
    username: usernameReducer,
});