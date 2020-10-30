package market

import (
	"fmt"
	"time"
)

type GameInstance struct {
	Running       bool
	ID            int
	TickRate      time.Duration
	TickTimestamp time.Time
	Market        *Market
}

func (game *GameInstance) Tick() {
	game.TickTimestamp = time.Now()
	game.Market.Update()
}

func (game GameInstance) String() string {
	return fmt.Sprintf("Game-%d: %v\n", game.ID, game.Market.String())
}
