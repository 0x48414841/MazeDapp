import maze from "../apis/maze";

//Action Creators
export const fetchMaze = () => async dispatch => {
    const response = await maze.get('/maze');
    dispatch({ type: 'FETCH_MAZE', payload: response })
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
