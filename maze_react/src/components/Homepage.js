import React from 'react';

const Homepage = () => {

    return (
        <div className="ui centered grid">
            <div className="four wide column center aligned ">
                <img className="ui medium circular image"
                    src="https://www.worldatlas.com/r/w1200-q80/upload/3f/3a/7e/shutterstock-529749241.jpg" />
            </div>
            <div className="eight wide column center aligned">
                <div>
                    <button className="ui button">
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