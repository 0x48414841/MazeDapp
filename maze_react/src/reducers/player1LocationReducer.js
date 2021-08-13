export default (state = [{ X: 0, Y: 0}], action) => {
    switch (action.type) {
        case 'UPDATE_POS':
            return action.payload;
        default:
            return state;
    }
};
