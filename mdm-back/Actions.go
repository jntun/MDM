package main

import (
	"fmt"

	"github.com/Nastyyy/mdm-back/market"
)

type Action interface {
	// Attempts to
	DoAction(msg *Message, sess *market.Session) (bool, error)
}

type BuyAction struct {
}

func (act *BuyAction) DoAction(msg *Message, sess *market.Session) (bool, error) {
	return true, nil
}

type SellAction struct {
}

func (act *SellAction) DoAction(msg *Message, sess *market.Session) (bool, error) {

	return true, nil
}

type PingAction struct {
}

func (act *PingAction) DoAction(msg *Message, sess *market.Session) (bool, error) {
	fmt.Println("Ping")
	fmt.Println(sess)

	return true, nil
}
