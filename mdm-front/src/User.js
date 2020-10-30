import React from 'react';
import Cookie from 'js-cookie'

export default class User extends React.Component {
  constructor(props) {
    super(props)

    this.state = {username: ''}
  }

  getUsername() {
    this.props.socket.sendData('test', {test: true})

    // TODO 
    return 'test'
  }

  componentDidMount() {
    console.log(this.getUsername())
  }

  updateName = (e) => {
    console.log("Updated name to:", this.state.username)
    Cookie.set('username', this.state.username)
  }

  onChange = (e) => {
    this.setState({username: e.target.value})
  }

  render() {
    return (
      <div id="username-input">
        <input value={this.state.username} onChange={this.onChange} className="userInput" placeholder="Username"/>
        <button onClick={this.updateName}>Update</button>
      </div>
    )
  }
}
