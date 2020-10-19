package main

import "fmt"

type Action interface {
	// Attempts to
	DoAction() (bool, error)
}

type BuyAction struct {
}

func (act *BuyAction) DoAction() {
}

type SellAction struct {
}

func (act *SellAction) DoAction() {
}

type PingAction struct {
}

func (act *PingAction) DoAction() {
	fmt.Println("Ping")
}
