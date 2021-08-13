//Action creators

export const setMaze = (maze) => ( { type: 'SET_MAZE', payload: { maze } });
export const wsConnect = url => ({ type: 'WS_CONNECT', payload: { url } });
export const wsConnecting = host => ({ type: 'WS_CONNECTING', payload: host });
export const wsConnected = host => ({ type: 'WS_CONNECTED', payload: host });
export const wsDisconnect = host => ({ type: 'WS_DISCONNECT', payload: host });
export const wsDisconnected = host => ({ type: 'WS_DISCONNECTED', payload: host });
export const setUsername = username => ({ type: 'SET_USERNAME', payload: username });

//make sure users moves to a valid square before sending request to server
export const updatePosFromClient = (currentCell, event) => {
    if (event.key === 'w' && currentCell.IsNLinked === true) {
        return { type: 'DEC_X' };
    } else if (event.key === 'a' && currentCell.IsWLinked === true) {
        return { type: 'DEC_Y' };
    } else if (event.key === 's' && currentCell.IsSLinked === true) {
        return { type: 'INC_X' };
    } else if (event.key === 'd' && currentCell.IsELinked === true) {
        return { type: 'INC_Y' };
    }
    return { type: 'Nop' }
};

export const updatePosFromServer = (allPos) =>  ({ type: 'UPDATE_POS', payload: allPos });
