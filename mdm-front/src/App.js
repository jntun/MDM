import React from 'react';
import Socket from './Socket';
import Cookie from 'js-cookie';
import axios from 'axios'
import './App.css';

export default class App extends React.Component {
  constructor(props) {
    super(props)
    const socket = new Socket("ws://"+this.props.endpoint);
    console.log(socket)
    this.socket = socket
    if(Cookie.get('uuid') == null) {
      this.setCookie()
    }
    this.socket.onopen = () => {
      this.socket.sendData({ping: true})
    }
  }

  componentDidMount() {
  }

  setCookie() {
    var endpoint = "http://" + this.props.endpoint + '/authorize';
    console.log("No uuid found... \n Attempting to authorize with:", endpoint)
    axios.get(endpoint).then((resp) => {
      Cookie.set('uuid', resp.data) 
    })
  }

  handleClick = (e) => {
    this.socket.sendData("PING", {"body": null})
  }

  render() {
    return (
      <div className="App">
        <p>SOCKET: {this.socket.OPEN ? 'open' : 'closed'}</p>
        <p>UUID: {Cookie.get('uuid')}</p>
        <button onClick={this.handleClick}>Ping</button>
      </div>
    );
  }
}
