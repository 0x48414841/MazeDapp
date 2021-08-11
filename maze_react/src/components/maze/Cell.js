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

    renderPlayers() {
       // console.log('render players', this.props);
        return (
            this.props.playersLoc.map(({X, Y}) => {
                const {Row, Col} = this.props.data;
                return (
                    <div>
                        {X == Row && Y == Col ? <Square currentCell={this.props.data} /> : null }
                    </div>
                );
            })
        );
    }
    
    /// {this.renderPlayers()}
//{this.props.shouldDisplaySquare === true ? <Square currentCell={this.props.data} /> : null}
    render() {
        return (

            <div className={this.buildCell()}>
                {this.renderPlayers()}
            </div>
        );
    };

}


const mapStateToProps = (state, ownProps) => {
    //const shouldDisplaySquare = state.playersLoc.X === ownProps.data.Row &&
      //  state.playersLoc.Y === ownProps.data.Col;
    return { playersLoc: state.playersLoc };
};

export default connect(mapStateToProps)(Cell);