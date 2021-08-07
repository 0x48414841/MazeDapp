import React, { useEffect } from 'react';
import { updatePos } from '../../actions';
import { connect } from 'react-redux';
import './Square.css';

const Square = ({currentCell, updatePos}) => {

    useEffect(() => {
        const onKeyPress = (event) => {
            if (event.key === 'w' || event.key === 'a' || event.key === 's' || event.key === 'd') {
                updatePos(currentCell, event);
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
    return { p1Loc: state.player1Location };
}

export default connect(mapStateToProps, {
    updatePos: updatePos //updatePos action creator
})(Square);