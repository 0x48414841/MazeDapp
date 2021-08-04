import React, { useState, useEffect } from 'react';
import { connect } from 'react-redux';
import Square from './Square';
import './Cell.css';

//do this with hooks!
/* const Cell = ({data}, shouldDisplaySquare) => {
    console.log(data, shouldDisplaySquare)
    const buildCell = () => {
        const output = "cell" +
            (data.IsNLinked === true ? " notop" : ' top') +
            (data.IsSLinked === true ? " nobottom" : ' bottom') +
            (data.IsWLinked === true ? " noleft" : ' left') +
            (data.IsELinked === true ? " noright" : ' right');
        return output;
    };

    return (
        <div className={buildCell()}>
            {shouldDisplaySquare === true ? <Square /> : null}
        </div>
    );
}; */
class Cell extends React.Component {
    buildCell = () => {

        const output = "cell" +
            (this.props.data.IsNLinked === true ? " notop"    : ' top') +
            (this.props.data.IsSLinked === true ? " nobottom" : ' bottom') +
            (this.props.data.IsWLinked === true ? " noleft"   : ' left') +
            (this.props.data.IsELinked === true ? " noright"  : ' right');
        return output;
    };

    render() {
        return (

            <div className={this.buildCell()}>
                {this.props.shouldDisplaySquare === true ? <Square currentCell={this.props.data} /> : null}
            </div>
        );
    };

}


const mapStateToProps = (state, ownProps) => {
    const shouldDisplaySquare = state.player1Location.X === ownProps.data.Row &&
        state.player1Location.Y === ownProps.data.Col;
    return { shouldDisplaySquare };
};

export default connect(mapStateToProps)(Cell);