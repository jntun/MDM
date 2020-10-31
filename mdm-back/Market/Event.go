package market

import (
	"log"
	"os"

	"github.com/gorilla/websocket"
	ms "github.com/mitchellh/mapstructure"
)

// Event map
const (
	ping     = "PING"
	buy      = "BUY"
	sell     = "SELL"
	register = "REGISTER"
)

func MapEvent(sess *Session, conn *websocket.Conn, message *Message) {
	if os.Getenv("DEBUG") == "true" {
		//fmt.Println("MessageObj:", *message)
	}

	var action Action

	switch message.Action {
	case buy:
		action = BuyAction{UUID: message.UUID.String()}
	case sell:
		action = SellAction{}
	case ping:
		action = PingAction{}
	case register:
		action = RegisterAction{uuid: message.UUID.String(), conn: conn}
	}

	err := ms.Decode(message.Body, &action)
	if err != nil {
		log.Println(err)
	}

	err = action.DoAction(sess)
	if err != nil {
		log.Printf("Error mapping event: %v", err)
		return
	}

	sess.SyncState()
}
