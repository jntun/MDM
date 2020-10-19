package main

import "fmt"

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

func EventHandler(message *Message) {
	fmt.Println("MessageObj:", *message)
	switch message.Action {
	// TODO: Fulfill events
	case buy:
		action := BuyAction{}
		action.DoAction()
	case sell:
		action := SellAction{}
		action.DoAction()
	case ping:
		action := PingAction{}
		action.DoAction()
	}
}
