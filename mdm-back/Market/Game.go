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
	TickTotal     uint64       `json:"tick"`
	Ticker        *time.Ticker `json:"-"`
	TickTimestamp time.Time    `json:"timestamp"`
	Market        *Market
}

func (game *GameInstance) Start() {
	game.Running = true
	if config.DEBUG && game.TickTotal == 0 {
		log.Printf("Starting game instance: %d\n", game.ID)
	} else {
		config.MainLog(fmt.Sprintf("Restarting game instance: %d", game.ID))
	}

	game.Run()
}

func (game *GameInstance) Stop() {
	game.Running = false
	game.Ticker.Stop()
	if config.DEBUG {
		log.Printf("Pausing game instance: %d\n", game.ID)
	}
}

func (game *GameInstance) Tick() {
	game.TickTotal++
	game.TickTimestamp = time.Now()
	game.Market.Update(game.TickTotal)
}

func (game *GameInstance) Run() {
	for range game.Ticker.C {
		startGameTime := time.Now()
		if game.Running {
			game.Tick()
		} else {
			config.VerboseLog(fmt.Sprintf("[Game-%d] Skipping game tick while paused...", game.ID))
		}

		endGameTime := time.Now()
		finalTime := endGameTime.Sub(startGameTime)

		config.VerboseLog(fmt.Sprintf("Tick: %d", game.TickTotal))
		config.PerfLog(fmt.Sprintf("[Game-%d] Tick took %v", game.ID, finalTime))
	}
}

func (game GameInstance) String() string {
	return fmt.Sprintf("Game-%d: %v\n", game.ID, game.Market.String())
}
