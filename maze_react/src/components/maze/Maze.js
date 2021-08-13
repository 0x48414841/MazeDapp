import React, { useEffect } from 'react';
import { connect, useSelector } from 'react-redux';
import { updatePosFromClient } from '../../actions';
import { useHistory } from 'react-router';
import { wsConnect } from '../../actions';
import Cell from './Cell';
import './Maze.css';

const Maze = ({ wsConnect, updatePosFromClient }) => {
    const history = useHistory();
    const maze = useSelector(store => store.maze);
    const playersLoc = useSelector(store => store.playersLoc);
    const username = useSelector(store => store.username);

    useEffect(() => {
        const { Id, Addr } = history.location.state;
        wsConnect(`ws://localhost${Addr}/game?id=${Id}`);

    }, [])

    const getPlayerLoc = () => {
        const host = playersLoc.find(player => player.Username === username);
        const {X, Y} = host.Pos;
        return maze[X][Y];
    };
    useEffect(() => {
        const onKeyPress = (event) => {

            if (event.key === 'w' || event.key === 'a' || event.key === 's' || event.key === 'd') {
                updatePosFromClient(getPlayerLoc(), event);
            }
        };

        document.body.addEventListener("keydown", onKeyPress, { capture: true });

        //enables graceful teardown
        return () => {
            document.body.removeEventListener("keydown", onKeyPress, {
                capture: true,
            });
        };
    }, [maze, playersLoc]);

    const renderMaze = () => {
        return maze.map(row => { //TODO add key for row
            return (
                <div className="row">
                    {
                        row.map(cell => {
                            return (
                                <React.Fragment key={`${cell.Row},${cell.Col}`}> <Cell data={cell} /> </React.Fragment>
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
                    <br />
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
    return { maze: state.maze, playersLoc: state.playersLoc };
}

export default connect(mapStateToProps, {
    wsConnect, updatePosFromClient 
})(Maze);