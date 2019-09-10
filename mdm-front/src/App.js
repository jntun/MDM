import React from 'react';
import Socket from './Socket'
import './App.css';

export default class App extends React.Component {
  constructor(props) {
    super(props)
    const socket = new Socket("ws://127.0.0.1:8080");
    console.log(socket)
    this.socket = socket

    this.state = {}
  }

  componentDidMount() {
    this.socket.sendData({data: "test data"});
  }

  render() {
    return (
      <div className="App">

      </div>
    );
  }
}