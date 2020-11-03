import React from 'react';

export default class Portfolio extends React.Component {

  MapHolding = (holding) => {
    return (
      <p key={holding.asset.ticker}>{holding.asset.ticker} | Shares: {holding.volume}</p>
    )
  }

  render() {
    var holdings = this.props.data;
    var displayHoldings = [];
    for(var ticker in holdings) {
      displayHoldings.push(this.MapHolding(holdings[ticker]))
    }



    return(
      <div id="portfolio">
        <h2>Portfolio</h2>
        {displayHoldings}
      </div>
    )
  }

}
