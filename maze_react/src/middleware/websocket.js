import { wsConnect, wsConnecting, wsConnected, wsDisconnect, wsDisconnected, setMaze, updatePosFromServer, setUsername } from '../actions'

const socketMiddleware = () => {
    let socket = null;

    const onOpen = store => (event) => {
        console.log('websocket open', event.target.url);
        store.dispatch(wsConnected(event.target.url));
    };

    const onClose = store => () => {
        store.dispatch(wsDisconnected());
    };

    const onMessage = store => (event) => {
        const payload = JSON.parse(event.data);

        //Receiving messages
        switch (payload.Action) {
            case 'RECV_MAZE':
                console.log('setting maze')
                store.dispatch(setMaze(payload.Maze));
                store.dispatch(setUsername(payload.Username));
                break;
            case 'RECV_POS':
                console.log(payload)
                store.dispatch(updatePosFromServer(payload.AllPos));
                break;
            default:
                break;
        }
    };

    // the middleware part of this function
    // send request to server
    return store => next => action => {
        switch (action.type) {
            case 'WS_CONNECT':
                console.log("sending request")
                if (socket !== null) {
                    socket.close();
                }

                // connect to the remote host
                socket = new WebSocket(action.payload.url);

                // websocket handlers
                socket.onmessage = onMessage(store);
                socket.onclose = onClose(store);
                socket.onopen = onOpen(store);

                break;
            case 'DEC_X':
                socket.send(JSON.stringify({ Action: "UP" }));
                break;
            case 'INC_X':
                socket.send(JSON.stringify({ Action: "DOWN" }));
                break;
            case 'DEC_Y':
                socket.send(JSON.stringify({ Action: "LEFT" }));
                break;
            case 'INC_Y':
                socket.send(JSON.stringify({ Action: "RIGHT" }));
                break;
            case 'WS_DISCONNECT':
                if (socket !== null) {
                    socket.close();
                }
                socket = null;
                console.log('websocket closed');
                break;
            case 'NEW_MESSAGE':
                console.log('sending a message', action.msg);
                socket.send(JSON.stringify({ command: 'NEW_MESSAGE', message: action.msg }));
                break;
            default:
                return next(action);
        }
    };
};

export default socketMiddleware();