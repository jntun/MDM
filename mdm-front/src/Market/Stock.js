import React from 'react';

export default class Stock extends React.Component {
  constructor(props) {
    super(props)

    this.state = {amount: 0}
  }

  buyAction = (e, volume) => {
    this.props.socket.sendData("BUY", {ticker: this.props.ticker, volume: volume})
  }

  render() {
    return (
      <div id={this.props.ticker} className="stock-container">
        <h1>{this.props.ticker}</h1>
        <p>${this.props.price} | {this.props.volume}</p>
        <button id={this.props.ticker + "-1"} onClick={(e) => this.buyAction(e, 1)}>Buy 1</button>
        <button id={this.props.ticker + "-5"} onClick={(e) => this.buyAction(e, 5)}>Buy 5</button>
        <button id={this.props.ticker + "-10"} onClick={(e) => this.buyAction(e, 10)}>Buy 10</button>
        <button id={this.props.ticker + "-20"} onClick={(e) => this.buyAction(e, 20)}>Buy 20</button>
      </div>
    )
  }
}
