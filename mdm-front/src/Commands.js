import React from 'react';

export default class Commands extends React.Component {
  constructor(props) {
    super(props)
  }

  Ping = (e) => {
    this.props.socket.sendData("PING", null)
  }


  Register = (e) => {
    this.props.socket.sendData("REGISTER", {})
  }

  render() {
    return (
      <div id="commands">
          <button onClick={this.Ping}>Ping</button><br/>
          <button onClick={this.Register}>Register</button>
      </div>
    )
  }
}
