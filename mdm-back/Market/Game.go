package market

import (
	"fmt"
	"log"
	"time"

	"github.com/Nastyyy/mdm-back/config"
)

type GameInstance struct {
	Running       bool
	ID            int
	Ticker        *time.Ticker `json:"-"`
	TickTimestamp time.Time
	Market        *Market
}

func (game *GameInstance) Start() {
	game.Running = true
	if config.DEBUG {
		log.Printf("Starting game instance: %d", game.ID)
	}
}

func (game *GameInstance) Stop() {
	game.Running = false
	if config.DEBUG {
		log.Printf("Pausing game instance: %d", game.ID)
	}
}

func (game *GameInstance) Tick() {
	game.TickTimestamp = time.Now()
	game.Market.Update()
}

func (game GameInstance) String() string {
	return fmt.Sprintf("Game-%d: %v\n", game.ID, game.Market.String())
}
