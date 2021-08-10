import React, { useState, useEffect } from 'react';
import { connect } from 'react-redux';
import { fetchMaze } from '../../actions';
import { useHistory } from 'react-router';
import { wsConnect } from '../../actions';
import Cell from './Cell';
import './Maze.css';

const Maze = ({ maze, wsConnect }) => {
    const history = useHistory();

    useEffect(() => {
        const { Id, Addr } = history.location.state;
        //wsConnect("ws://localhost:8909/");

        wsConnect(`ws://localhost${Addr}/game?id=${Id}`);
    }, [])

    const renderMaze = () => {
        return maze.map(row => {
            return (
                <div className="row">
                    {
                        row.map(cell => {
                            return (
                                //<div class="cell left top right bottom"></div>
                                <React.Fragment> <Cell data={cell} /> </React.Fragment>
                            );
                        })
                    }
                </div>
            );
        });
    };

    return (
        <div className="ui grid">
            <div className="three wide column">
                <div className="ui container center aligned">
                    <h1> MAKE SEPARATE COMP Player 1 </h1>
                </div>
            </div>
            <div className="ten wide column">
                <div className="ui container center aligned">
                    <div> {renderMaze()}</div>
                    <br/>
                    <div>
                        <h2>MAKE SEPARATE COMP  Wager</h2>

                        <button class="ui labeled icon button">
                            <i class="arrow down icon"></i>
                            - 1
                        </button>
                        <div class="ui input">
                            <input type="text" placeholder="Bet..." />
                        </div>
                        <button class="ui labeled icon button">
                            <i class="arrow up icon"></i>
                            + 1
                        </button>
                    </div>
                </div>

            </div>
            <div className="three wide column">
                <div className="ui container center aligned">
                    <h1> MAKE SEPARATE COMP  Player 2 </h1>
                </div>
            </div>

        </div>

    );
};

const mapStateToProps = (state) => {
    return { maze: state.maze };
}

export default connect(mapStateToProps, {
    wsConnect, //fetchMaze action creator
})(Maze);