import React, { useEffect } from 'react';
import { updatePosFromClient } from '../../actions';
import { connect } from 'react-redux';
import './Square.css';

const Square = ({currentCell, updatePosFromClient}) => {

    useEffect(() => {
        const onKeyPress = (event) => {

            if (event.key === 'w' || event.key === 'a' || event.key === 's' || event.key === 'd') {
                updatePosFromClient(currentCell, event);
            }
        };

        document.body.addEventListener("keydown", onKeyPress, { capture: true });
        return () => {
            document.body.removeEventListener("keydown", onKeyPress, {
                capture: true,
            });
        };
    }, []);

    return <div className="square"></div>;
};

const mapStateToProps = (state) => {
    return { p1Loc: state.playersLoc};
}

export default connect(mapStateToProps, {
    updatePosFromClient //updatePos action creator
})(Square);