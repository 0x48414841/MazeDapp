import { combineReducers } from "redux";
import mazeReducer  from "./mazeReducer";
import player1LocationReducer from "./player1LocationReducer";

export default combineReducers({
    maze: mazeReducer,
    player1Location: player1LocationReducer,
});