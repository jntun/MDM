package market

import (
	"log"
	"os"
)

// Event map
const (
	ping = "PING"
	buy  = "BUY"
	sell = "SELL"
)

func MapEvent(sess *Session, message *Message) {
	if os.Getenv("DEBUG") == "true" {
		//fmt.Println("MessageObj:", *message)
	}
	var action Action

	/*
		-- Probably don't need this now --
		user, err := sess.GetUser(message.UUID.String())
		if err != nil {
			log.Printf("Error getting user: %v", err)
		}
	*/

	switch message.Action {
	// TODO: Fulfill events
	case buy:
		action = BuyAction{
			UUID:   message.UUID.String(),
			Ticker: message.Body.(map[string]interface{})["ticker"].(string),
			Volume: int(message.Body.(map[string]interface{})["volume"].(float64)),
		}
	case sell:
		action = SellAction{}
	case ping:
		action = PingAction{}
	}

	err := action.DoAction(sess)
	if err != nil {
		log.Printf("Error mapping event: %v", err)
	}

	sess.SyncState()
}
