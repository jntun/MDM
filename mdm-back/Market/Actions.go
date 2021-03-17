package market

import (
	"fmt"

        "github.com/Nastyyy/mdm-back/config"
	"github.com/gorilla/websocket"
)

type Action interface {
    // Any struct that has a DoAction with *Session and *User 
    // as parameters fits the criteria for being a game action. 
    // error is subject to the action being performed.
    DoAction(sess *Session, usr *User) error
}

type BuyAction struct {
	UUID   string
	Ticker string
	Volume int
}

func (buy BuyAction) DoAction(sess *Session, usr *User) error {
	user := usr
	if user == nil {
		return fmt.Errorf("[BuyAction] %s - not found in session", buy.UUID)
	}

	stock, err := sess.Game.Market.GetStock(buy.Ticker)
	if err != nil {
		return err
	}

	if !stock.CanBuy(buy.Volume, user.GetBalance()) {
		return fmt.Errorf("[BuyAction] %s cannot buy %d %s", user.Name, buy.Volume, buy.Ticker)
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

	config.MainLog(fmt.Sprintf("[BuyAction] %s bought %s for %v | Balance: %v | %v\n", user.Name, stock.Ticker, cost, user.GetBalance(), stock))
	return nil
}

func (buy BuyAction) String() string {
	return "BuyAction: " + buy.Ticker
}

type SellAction struct {
	UUID   string
	Ticker string
	Volume int
}

func (sell SellAction) DoAction(sess *Session, usr *User) error {
	user := sess.GetUser(sell.UUID)
	if user == nil {
		return fmt.Errorf("%s - not found in session", sell.UUID)
	}

	stock, err := sess.Game.Market.GetStock(sell.Ticker)
	if err != nil {
		return err
	}

	err = user.CanSellHolding(stock, sell.Volume)
	if err != nil {
		return err
	}

	user.UpdateHolding(stock, -sell.Volume)

	amount := (stock.Price * float32(sell.Volume))
	user.Deposit(amount)

	config.MainLog(fmt.Sprintf("%s sold: %s for %v | Balance: %v | %v\n", user.Name, stock.Ticker, amount, user.GetBalance(), stock))
	return nil
}

type PingAction struct {
}

func (act PingAction) DoAction(sess *Session, usr *User) error {
	config.DebugLog(fmt.Sprintf("PingAction: sess: %s\n", *sess))
	return nil
}

type RegisterAction struct {
	UUID string
	Name string `json:"name,omitempty"`
	conn *websocket.Conn
}

func (reg RegisterAction) DoAction(sess *Session, usr *User) error {

	// Check if user is already in session
	user := sess.GetUser(reg.UUID)
        //user = usr
	if user != nil {
		fmt.Printf("User: %s found in session, updating connection...\n", user.Name)
		user.Conn = reg.conn
		return nil
	}

	// If not make them
	// sess.NewUser(reg.
	user, err := NewUser(reg.Name, reg.UUID)
	if err != nil {
		return err
	}
	config.MainLog(fmt.Sprintf("[RegisterAction] %s", user))
	user.Conn = reg.conn

	sess.AddUser(user)
	return nil
}

type UsernameAction struct {
	Username string `json:"username"`
}

func (act UsernameAction) DoAction(sess *Session, usr *User) error {
	user := usr
	if act.Username == "" {
            return fmt.Errorf("[UsernameAction] Provided empty username for: %s", usr)
	}

	config.MainLog(fmt.Sprintf("[UsernameAction] %s changed name to %s...\n", user.Name, act.Username))
	user.Name = act.Username
	return nil
}
