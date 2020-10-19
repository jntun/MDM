package market

import (
	"time"
)

type GameInstance struct {
	Running  bool
	TickRate time.Duration
	Market   *Market
}

func (game *GameInstance) Tick() {
	game.Market.Update()
}

func (game GameInstance) String() string {
	return game.Market.String()
}
