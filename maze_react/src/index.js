import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import { createStore, applyMiddleware } from 'redux';
import thunk from 'redux-thunk';
import wsMiddleware from './middleware/websocket';

import App from './components/App';
import reducers from './reducers';


const store = createStore(reducers, applyMiddleware(thunk, wsMiddleware));

ReactDOM.render(
    <Provider store={store}>
        <App />
    </Provider>,
    document.querySelector('#root'),
);