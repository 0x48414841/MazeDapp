import maze from "../apis/maze";

//Action Creators
export const fetchMaze = () => async dispatch => {
    const response = await maze.get('/maze');
    dispatch({ type: 'FETCH_MAZE', payload: response })
    console.log(response)
    dispatch(updateMaxRowAndCol(response.data.Maze.length, response.data.Maze[0].length ));
};

export const updatePos = (currentCell, event) => {
    if (event.key === 'w' && currentCell.IsNLinked === true) {
        return { type: 'DEC_X' };
    } else if (event.key === 'a'  && currentCell.IsWLinked === true) {
        return { type: 'DEC_Y' };
    } else if (event.key === 's'  && currentCell.IsSLinked === true) {
        return { type: 'INC_X' };
    } else if (event.key === 'd'  && currentCell.IsELinked === true) {
        return { type: 'INC_Y' };
    }
    return {type: 'Nop'}
};

export const updateMaxRowAndCol = (LastRow, LastCol) => {
   return {type: 'UPDATE_MAX_ROW_AND_COL', payload: {LastRow, LastCol}};
};
