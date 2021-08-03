import maze from "../apis/maze";

//Action Creators
export const fetchMaze = () => async dispatch => {
    const response = await maze.get('/maze');
    dispatch({ type: 'FETCH_MAZE', payload: response})
};