import Cookies from 'js-cookies';

export default class Socket {
  constructor(ip, onmessage) {
    this.ip = ip;
    this.onmessage=onmessage;
    this.open();
    this.data = [];
  }

  open = () => {
    this.socket = new WebSocket(this.ip);
    this.socket.onmessage = (e) => {
      this.data.push(e);
      this.onmessage(e);
    }
  }

  sendData = (action, data) => {
      try {
          var sendData = JSON.stringify({
              uuid: Cookies.getItem('uuid'),
              action: action,
              body: data
          });
          this.socket.send(sendData);
          console.log(sendData);
        return true;
      } catch(err) {
          console.log(err);
        return false;
      }
  }

  /*
  sleep = async (ms) => {
    return new Promise(resolve => setTimeout(resolve, ms));
  }
  */
}
