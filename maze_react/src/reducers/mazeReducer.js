export default (state = [[]], action) => {
    switch (action.type) {
        case 'SET_MAZE':
            return action.payload.maze;
        default:
            return state;
    }
};