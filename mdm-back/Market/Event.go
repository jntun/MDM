package market

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	ms "github.com/mitchellh/mapstructure"
)

// Event map
const (
	PING       = "PING"
	BUY        = "BUY"
	SELL       = "SELL"
	REGISTER   = "REGISTER"
	UPDATENAME = "USERNAME"
        ADMIN      = "ADMIN"
)

func MapEvent(sess *Session, conn *websocket.Conn, message *Message) {
	var action Action

	switch message.Action {
	case BUY:
		action = BuyAction{UUID: message.UUID.String()}
	case SELL:
		action = SellAction{UUID: message.UUID.String()}
	case PING:
		action = PingAction{}
	case REGISTER:
		action = RegisterAction{UUID: message.UUID.String(), conn: conn}
	case UPDATENAME:
		action = UsernameAction{}
        case ADMIN:
                action = AdminAction{}

	default:
		fmt.Printf("Invalid event received: %v\n", message)
		return
	}

	err := ms.Decode(message.Body, &action)
	if err != nil {
		log.Println(err)
	}

        usr := sess.Users[message.UUID.String()]
        /* Not sure about how to do this as of right now. Will have to revisit.
        if usr == nil {
            log.Printf("[ERROR][MapEvent] Could not find user in session: %s\n", message.UUID.String())
            err = action.DoAction(sess, nil)
            if err != nil {
                log.Printf("[ERROR][MapEvent] Error performing action with nil usr: %s\n", err)
            }
        }
        */


        // TODO: Refactor all action DoAction methods
        // Refactor to action.DoAction(sess *Session, user *User) 
        // All actions are events handled from users so it only makes sense
        // that all actions are user context aware. Also it would lead to much cleaner
        // access flow so I'm not doing the uuid: message.UUID.String() over and over
        err = action.DoAction(sess, usr)
	if err != nil {
		log.Printf("[ERROR][MapEvent]%s\n", err)
		return
	}

	sess.SyncState()
}
