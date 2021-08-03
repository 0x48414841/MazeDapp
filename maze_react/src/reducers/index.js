import { combineReducers } from "redux";
import mazeReducer  from "./mazeReducer";

export default combineReducers({
    maze: mazeReducer,
});