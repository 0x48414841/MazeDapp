export default (state = [[]], action) => {
    switch (action.type) {
        case 'SET_MAZE':
            console.log(action)
            return action.payload.maze;
        default:
            return state;
    }
};