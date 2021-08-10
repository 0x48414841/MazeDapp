import React from 'react';
import { useHistory } from 'react-router';
import maze from "../apis/maze";
import axios from 'axios';

const Homepage = () => {
    const history = useHistory();
    const onCreateGameClick = () => {
        const getGameURL = async () => {
            return await axios.get('http://localhost:8000/createGame')
                .then(({data}) => {
                    const {Id, Addr} = data;
                    console.log(Id, Addr)
                    history.push({
                        pathname: '/game',
                        state: { Id, Addr }
                    });
                });

            //return await maze.get('http://localhost:8000/createGame').data;
        }


        getGameURL()
    }

    return (
        <div className="ui centered grid">
            <div className="four wide column center aligned ">
                <img className="ui medium circular image"
                    src="https://www.worldatlas.com/r/w1200-q80/upload/3f/3a/7e/shutterstock-529749241.jpg" />
            </div>
            <div className="eight wide column center aligned">
                <div>
                    <button className="ui button" onClick={onCreateGameClick}>
                        Create Game
                    </button>
                </div>
                <br />
                <div>
                    <button className="ui button">
                        Join Lobby
                    </button>
                </div>
            </div>
            <div className="four wide column">Current Lobbies</div>
        </div>
    );
};

export default Homepage;