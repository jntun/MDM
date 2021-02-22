package market

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	ms "github.com/mitchellh/mapstructure"
)

// Event map
const (
	ping       = "PING"
	buy        = "BUY"
	sell       = "SELL"
	register   = "REGISTER"
	updatename = "USERNAME"
        admin      = "ADMIN"
)

func MapEvent(sess *Session, conn *websocket.Conn, message *Message) {
	var action Action

	switch message.Action {
	case buy:
		action = BuyAction{UUID: message.UUID.String()}
	case sell:
		action = SellAction{UUID: message.UUID.String()}
	case ping:
		action = PingAction{}
	case register:
		action = RegisterAction{uuid: message.UUID.String(), conn: conn}
	case updatename:
		action = UsernameAction{uuid: message.UUID.String()}
        case admin:
                action = AdminAction{uuid: message.UUID.String()}
	default:
		fmt.Printf("Invalid event received: %v\n", message)
		return
	}

	err := ms.Decode(message.Body, &action)
	if err != nil {
		log.Println(err)
	}

        // TODO: Refactor all action DoAction methods
        // Refactor to action.DoAction(sess *Session, user *User) 
        // All actions are events handled from users so it only makes sense
        // that all actions are user context aware. Also it would lead to much cleaner
        // access flow so I'm noot doing the uuid: message.UUID.String() over and over
	err = action.DoAction(sess)
	if err != nil {
		log.Printf("[ERROR][MapEvent]%s\n", err)
		return
	}

	sess.SyncState()
}
