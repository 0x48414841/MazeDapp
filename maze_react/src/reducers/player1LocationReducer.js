export default (state = { X: 0, Y: 0, LastRow: 0, LastCol: 0 }, action) => {
    switch (action.type) {
        case 'INC_X':
            return { ...state, X: min(state.X + 1, state.LastRow), Y: state.Y };
        case 'DEC_X':
            return { ...state, X: max(state.X - 1, 0), Y: state.Y };
        case 'INC_Y':
            return { ...state, X: state.X, Y: min(state.Y + 1, state.LastCol) }
        case 'DEC_Y':
            return {...state,  X: state.X, Y: max(state.Y - 1, 0) };
        case 'UPDATE_MAX_ROW_AND_COL': //this allows for dfferent sized mazes
            return { ...state, LastRow: action.payload.LastRow, LastCol: action.payload.LastCol };
        default:
            return state;
    }
};

const min = (a, b) => {
    return a < b ? a : b;
}


const max = (a, b) => {
    return a > b ? a : b;
}