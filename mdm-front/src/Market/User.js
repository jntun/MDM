import React from 'react';
import Portfolio from './Portfolio.js';

export default class User extends React.Component {
  constructor(props) {
    super(props);

    var username;
    if(this.props.data.username !== null) {
      username = this.props.data.username;
    } else {
      username = 'loading...';
    }

    this.state = {username: username, editing: false}
  }

  toggleEditing = (e) => {
    this.setState({editing: this.state.editing ? false : true});
    console.log(this.state.editing);
  }

  setUsername = (e) => {
    this.setState({username: e.target.value});
  }

  updateUsername = (e) => {
    this.props.socket.UpdateUsername(this.state.username);
    this.toggleEditing(e)
  }


  render() {
    var portfolio;
    portfolio = <Portfolio socket={this.socket} data={this.props.data.portfolio}/>

    var username;
    if(this.state.editing) {
      username = <input autoFocus={true} onBlur={this.updateUsername}onChange={this.setUsername} placeholder={this.props.data.username}/>
    } else {
      username = <h1 onClick={this.toggleEditing}>{this.props.data.username}</h1>
    }

    return (
      <div id="user">
        {username}
        <p>Balance: ${this.props.data.balance}</p>
        {portfolio}
      </div>
    )
  }
}
