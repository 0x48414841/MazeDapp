import React, { useEffect } from 'react';
import { updatePosFromClient } from '../../actions';
import { connect } from 'react-redux';
import './Square.css';

const Square = ({currentCell, updatePosFromClient}) => {

    useEffect(() => {
        
    }, []);

    return <div className="square"></div>;
};

const mapStateToProps = (state) => {
    return { p1Loc: state.playersLoc};
}

export default connect(mapStateToProps, {
    updatePosFromClient //updatePos action creator
})(Square);