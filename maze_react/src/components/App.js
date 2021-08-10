import React from 'react';
import { BrowserRouter, Route } from 'react-router-dom';
import Maze from './maze/Maze';
import Homepage from './Homepage';

const App = () => {
    return (
        <div className="ui container" >
            <BrowserRouter>
                <div>
                    <Route path="/" exact component={Homepage}/>
                    <Route path="/game" component={Maze}/>
                </div>
            </BrowserRouter>
        </div>
    );
};

export default App;