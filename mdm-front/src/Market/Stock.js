import React from 'react';

export default class Stock extends React.Component {
  constructor(props) {
    super(props)

    this.state = {amount: 0}
  }

  buyAction = (e, volume) => {
    //this.props.socket.sendData("BUY", {ticker: this.props.ticker, volume: volume})
    this.props.socket.Buy(this.props.ticker, volume);
  }

  sellAction = (e, volume) => {
    this.props.socket.Sell(this.props.ticker, volume);
  }

  render() {
    return (
      <div id={this.props.ticker} className="stock-container">
        <h1>{this.props.ticker}</h1>
        <p>${this.props.price} | Volume: {this.props.volume}</p>
        <button className="stock-button" id={this.props.ticker + "-1"} onClick={(e) => this.buyAction(e, 1)}>Buy 1</button>
        <button className="stock-button" id={this.props.ticker + "-5"} onClick={(e) => this.buyAction(e, 5)}>Buy 5</button>
        <button className="stock-button" id={this.props.ticker + "-10"} onClick={(e) => this.buyAction(e, 10)}>Buy 10</button>
        <br/>
        <button className="stock-button" id={this.props.ticker + "-1"} onClick={(e) => this.sellAction(e, 1)}>Sell 1</button>
        <button className="stock-button" id={this.props.ticker + "-5"} onClick={(e) => this.sellAction(e, 5)}>Sell 5</button>
        <button className="stock-button" id={this.props.ticker + "-10"} onClick={(e) => this.sellAction(e, 10)}>Sell 10</button>

      </div>
    )
  }
}
