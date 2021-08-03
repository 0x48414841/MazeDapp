import React, { useState, useEffect } from 'react';
import { connect } from 'react-redux';
import { fetchMaze } from '../actions';
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
                                <div class="cell left top right bottom"></div>
                            );
                        })
                    }
                </div>
            );
        });
    };

    return (
        <div>{renderMaze()}</div>
    );
};

const mapStateToProps = (state) => {
    return { maze: state.maze };
}

export default connect(mapStateToProps, {
    fetchMaze //fetchMaze action creator
})(Maze);