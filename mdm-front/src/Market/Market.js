import React from 'react';
import Stock from './Stock.js';

export default class Market extends React.Component {
  constructor(props) {
    super(props)
  }

  getStocks = () => {
    var stocks = this.props.marketData.Stocks;
    stocks = stocks.map(stock => {
      return <Stock socket={this.props.socket} key={stock.ticker} ticker={stock.ticker} price={stock.price} volume={stock.volume}/>;
    })
    return stocks;
  }

  render() {
    var stocks = null
    try {
      //console.log("Render():", this.getStocks())
      stocks = this.getStocks();
      } catch(e) {
      //throw e;
    }

    return (
      <div id="market-container">

        <button onClick={this.test}>Click me</button>
      {stocks}
      </div>
    );
  }
}
