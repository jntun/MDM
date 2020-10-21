package main

import (
	"fmt"

	"github.com/Nastyyy/mdm-back/market"
)

// Event map
const (
	ping = "PING"
	buy  = "BUY"
	sell = "SELL"
)

/*
* Not sure if needed
type Event struct {
}
*/

func EventHandler(message *Message, session *market.Session) {
	fmt.Println("MessageObj:", *message)
	switch message.Action {
	// TODO: Fulfill events
	case buy:
		action := BuyAction{}
		action.DoAction(message, session)
	case sell:
		action := SellAction{}
		action.DoAction(message, session)
	case ping:
		action := PingAction{}
		action.DoAction(message, session)
	}
}
