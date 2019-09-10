import Cookies from 'js-cookies';

export default class Socket {
    constructor(ip) {
        this.ip = ip;
        this.open();
    }

    open = () => {
        this.socket = new WebSocket(this.ip);
        this.socket.open
    }

    sendData = (data) => {
        try {
            var sendData = JSON.stringify({
                uuid: Cookies.getItem('uuid'),
                body: data
            });
            console.log(sendData);
            this.socket.send(sendData);
            console.log(sendData);
        } catch(err) {
            console.log(err);
        }
    }
}