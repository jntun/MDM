package market

import (
	"errors"
	"fmt"

	"github.com/gorilla/websocket"
)

type Action interface {
	// Attempts to
	DoAction(sess *Session) error
}

type BuyAction struct {
	UUID   string `json:"uuid"`
	Ticker string
	Volume int
}

func (buy BuyAction) DoAction(sess *Session) error {
	user, err := sess.GetUser(buy.UUID)
	if err != nil {
		return err
	}

	stock, err := sess.Game.Market.GetStock(buy.Ticker)
	if err != nil {
		return err
	}

	if !stock.CanBuy(buy.Volume, user.GetBalance()) {
		return errors.New(fmt.Sprintf("BuyAction: %s cannot buy %d %s", user.Name, buy.Volume, buy.Ticker))
	}

	err = user.UpdateHolding(stock, buy.Volume)
	if err != nil {
		return err
	}

	cost := stock.Price * float32(buy.Volume)
	err = user.Withdraw(cost)
	if err != nil {
		return err
	}

	fmt.Printf("%s bought: %s for %v | Balance: %v\n", user.Name, stock.Ticker, cost, user.GetBalance())
	return nil
}

func (buy BuyAction) String() string {
	return "BuyAction: " + buy.Ticker
}

type SellAction struct {
}

func (act SellAction) DoAction(sess *Session) error {
	return nil
}

type PingAction struct {
}

func (act PingAction) DoAction(sess *Session) error {
	fmt.Printf("PingAction: sess: %s\n", *sess)
	return nil
}

type RegisterAction struct {
	uuid string
	conn *websocket.Conn
}

func (reg RegisterAction) DoAction(sess *Session) error {
	// Check for if user is already in session
	user, err := sess.GetUser(reg.uuid)
	if err == nil {
		user.Conn = reg.conn
		return nil
	}
	// If not make them
	// sess.NewUser(reg.
	return nil
}
