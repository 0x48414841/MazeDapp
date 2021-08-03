export default (state = [[]], action) => {
    switch (action.type) {
        case 'FETCH_MAZE':
            console.log(action)
            return action.payload.data.Maze;
        default:
            return state;
    }
};