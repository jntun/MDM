package market

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Nastyyy/mdm-back/config"
	"github.com/Pallinder/go-randomdata"
	"github.com/gorilla/websocket"
	satori "github.com/satori/go.uuid"
)

var LIMIT float32 = 1000000000.0

type User struct {
	Name      string              `json:"username"`
	Timestamp *time.Time          `json:"timestamp"`
	UUID      *satori.UUID        `json:"uuid"`
	Portfolio map[string]*Holding `json:"portfolio"`
	Balance   float32             `json:"balance"`
	Conn      *websocket.Conn     `json:"-"`
	Mu        sync.Mutex
}

type Holding struct {
	Asset  *Stock `json:"asset"`
	Volume int    `json:"volume"`
}

func (user *User) Withdraw(amount float32) error {
	if user.Balance-amount >= 0.0 {
		user.Balance -= amount
		return nil
	}
	return fmt.Errorf("Error withdrawing from %v: %v", user, amount)
}

func (user *User) Deposit(amount float32) error {
	if user.Balance+amount <= LIMIT {
		user.Balance += amount
		return nil
	}

	return fmt.Errorf("Error depositing into %v: %v", user, amount)
}

func (user *User) UpdateHolding(stock *Stock, volume int) error {
	holding := user.Portfolio[stock.Ticker]
	if holding == nil {
		holding := &Holding{stock, volume}
		user.Portfolio[stock.Ticker] = holding
		return nil
	}
	/*
		if holding.Volume != 0 || holding.Volume-volume >= 0 {
			holding.Volume += volume
			return nil
		}
	*/
	holding.Volume += volume
	return nil
	//return fmt.Errorf("%s can't update holding: %s | volume: %d", user.Name, stock.Ticker, volume)
}

// SendUpdate is how a session sends itself to the user struct
// TODO: Hijack the payload here at the user level and determine what state to strip so the user
// 	 only gets state that they *need*.
func (user *User) SendUpdate(sess *Session) error {
	if user.Conn != nil {

		if config.DEBUG_VERBOSE {
			log.Printf("[%s] writing session...", user.Name)
		}

		if err := user.Conn.WriteJSON(sess); err != nil {
			user.Conn = nil
			return err
		}
		return nil
	}
	return fmt.Errorf("[%s] doesn't have active connection", user.Name)
}

func (user *User) CanSellHolding(stock *Stock, volume int) error {
	holding := user.Portfolio[stock.Ticker]
	if holding == nil {
            return fmt.Errorf("[%s] cannot sell %s: holding doesn't exist", user.Name, stock.Ticker)
	}

	if volume > holding.Volume {
            return fmt.Errorf("[%s] cannot sell %s: sell volume exceeds holding volume", user.Name, stock.Ticker)
	}

	return nil
}

func (user *User) SetUUID(newUUID string) error {
	gen, err := satori.FromString(newUUID)
	if err != nil {
		return err
	}

	user.UUID = &gen
	return nil
}

func (user *User) GetWorth() float32 {
	return user.Balance
}

func (user *User) GetPortfolioVolume(ticker string) int {
	for _, holding := range user.Portfolio {
		if holding.Asset.Ticker == ticker {
			return holding.Volume
		}
	}
	return 0
}

func (user *User) GetBalance() float32 {
	return user.Balance
}

func (user *User) String() string {
	return fmt.Sprintf("%s -- %v | Worth:$%v | Balance:$%v |", user.Name, user.UUID, user.GetWorth(), user.GetBalance())
}

// NewUser creats and initializes a new user instance
func NewUser(name string, uuid string) (*User, error) {
	currentTime := time.Now()
	playerUUID, err := satori.FromString(uuid)
	if err != nil {
		return nil, err
	}

	if name == "" {
		name = randomdata.State(randomdata.Large)
	}

	portfolio := make(map[string]*Holding)
	player := User{Name: name, Timestamp: &currentTime, UUID: &playerUUID, Portfolio: portfolio, Balance: 100000}
	return &player, nil
}
