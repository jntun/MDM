import React from 'react';
import Socket from './Socket';
import Commands from './Commands.js'
import Market from './Market/Market.js'
import Cookie from 'js-cookie';
import axios from 'axios'
import './App.css';

export default class App extends React.Component {
  constructor(props) {
    super(props)
    const socket = new Socket("ws://"+this.props.endpoint, this.update);
    console.log(socket)
    this.socket = socket
    if(Cookie.get('uuid') == null) {
      this.setCookie()
    }
    /*
    this.socket.onopen = () => {
      this.socket.sendData({ping: true})
    }
    this.socket.onmessage = (e) => {
      console.log(e);
    }
    */
    this.state = {stream: false, data: null}
  }

  componentDidMount() {
  }

  update = (e) => {
  console.log("--Message received--")
  try {
        var jsonData = JSON.parse(this.socket.data[this.socket.data.length-1]["data"]);
        this.setState({data: jsonData});
      } catch(e) {
        console.log(e)
      }
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
    var market;
    if(this.state.data !== null) {
      market = <Market socket={this.socket} marketData={this.state.data.game.Market}/>
    } else {
      market = <Market socket={this.socket} marketData={this.state.data}/>
    }

    return (
      <div className="App">
        <p>UUID: {Cookie.get('uuid')}</p>
        <Commands socket={this.socket}/>
        {market}
      </div>
    );
  }
}
