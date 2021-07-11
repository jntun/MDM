package market

import (
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

func MapEvent(sess *Session, conn *websocket.Conn, msg *Message) {
	var action Action

	switch msg.Action {
	case BUY:
		action = BuyAction{UUID: msg.UUID.String()}
	case SELL:
		action = SellAction{UUID: msg.UUID.String()}
	case PING:
		action = PingAction{}
	case REGISTER:
		action = RegisterAction{UUID: msg.UUID.String(), conn: conn}
	case UPDATENAME:
		action = UsernameAction{}
	case ADMIN:
		action = AdminAction{msg.Body}
	default:
		log.Printf("[ERROR] Invalid event received: %s\n", msg)
		return
	}

	err := ms.Decode(msg.Body, &action)
	if err != nil {
		log.Printf("error MapEvent(): \n%s\n\n", err)
		return
	}

	usr := sess.Users[msg.UUID.String()]
	/* Not sure about how to do this as of right now. Will have to revisit.
	   if usr == nil {
	       log.Printf("[ERROR][MapEvent] Could not find user in session: %s\n", msg.UUID.String())
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
	// access flow so I'm not doing the uuid: msg.UUID.String() over and over
	err = action.DoAction(sess, usr)
	if err != nil {
		log.Printf("[ERROR][MapEvent]%s\n", err)
		return
	}

	sess.SyncState()
}
