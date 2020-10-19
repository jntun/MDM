import Cookies from 'js-cookies';

export default class Socket {
  constructor(ip) {
    this.ip = ip;
    this.open();
  }

  open = () => {
    this.socket = new WebSocket(this.ip);
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
