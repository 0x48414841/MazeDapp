import React, { useState, useEffect } from 'react';
import { connect } from 'react-redux';
import { fetchMaze } from '../../actions';
import Cell from './Cell';
import './Maze.css';

const Maze = ({ maze, fetchMaze }) => {
    useEffect(() => {
        fetchMaze();
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
                    <h2>MAKE SEPARATE COMP  Wager</h2>
                    <div>
                        <button class="ui labeled icon button">
                            <i class="arrow down icon"></i>
                            - 1
                        </button>
                        <div class="ui input">
                            <input type="text" placeholder="Search..." />
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
    fetchMaze, //fetchMaze action creator
})(Maze);