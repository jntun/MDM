import React from 'react';
import Portfolio from './Portfolio.js';

export default class User extends React.Component {
  constructor(props) {
    super(props);

    this.state = {username: 'loading'}
  }

  render() {
    var portfolio;
    portfolio = <Portfolio socket={this.socket} data={this.props.data.portfolio}/>

    return (
      <div id="user">
        <h1>{this.props.data.username}</h1>
        <p>Balance: ${this.props.data.balance}</p>
        {portfolio}
      </div>
    )
  }
}
