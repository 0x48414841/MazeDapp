export default (state = { X: 0, Y: 0}, action) => {
    //Error checking was done in updatePos action creator
    switch (action.type) {
        case 'INC_X':
            return { ...state, X: state.X + 1};
        case 'DEC_X':
            return { ...state, X: state.X - 1};
        case 'INC_Y':
            return { ...state, Y: state.Y + 1 }
        case 'DEC_Y':
            return {...state, Y: state.Y - 1 };
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