package market

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Action interface {
	// Attempts to
	DoAction(sess *Session) error
}

type BuyAction struct {
	UUID   string
	Ticker string
	Volume int
}

func (buy BuyAction) DoAction(sess *Session) error {
	user := sess.GetUser(buy.UUID)
	if user == nil {
		return fmt.Errorf("%s - not found in session", buy.UUID)
	}

	stock, err := sess.Game.Market.GetStock(buy.Ticker)
	if err != nil {
		return err
	}

	if !stock.CanBuy(buy.Volume, user.GetBalance()) {
		return fmt.Errorf("BuyAction: %s cannot buy %d %s", user.Name, buy.Volume, buy.Ticker)
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
	stock.Volume -= buy.Volume

	fmt.Printf("%s bought: %s for %v | Balance: %v | %v\n", user.Name, stock.Ticker, cost, user.GetBalance(), stock)
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
	Name string `json:"name,omitempty"`
	conn *websocket.Conn
}

func (reg RegisterAction) DoAction(sess *Session) error {
	// Check for if user is already in session
	user := sess.GetUser(reg.uuid)
	if user != nil {
		fmt.Printf("User: %s found in session, updating connection...\n", user.Name)
		user.Conn = reg.conn
		return nil
	}
	// If not make them
	// sess.NewUser(reg.
	fmt.Printf("Creating new user object...")
	user, err := NewUser(reg.Name, reg.uuid)
	if err != nil {
		return err
	}
	user.Conn = reg.conn

	sess.AddUser(user)
	return nil
}
