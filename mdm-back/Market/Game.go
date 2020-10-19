package market

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type GameInstance struct {
	running bool
	Market  *Market
}

func (game *GameInstance) Tick() {
	game.Market.Update()
}

func (game GameInstance) String() string {
	return game.Market.String()
}

func Game() {
	duration := time.Minute * 2

	if os.Getenv("DEBUG") == "true" {
		duration = time.Millisecond * 500
	}

	game := GameInstance{running: true, Market: NewMarket()}
	for i := 0; game.running; i++ {
		//if os.Getenv("DEBUG") == "true" {
		printGame(&game, i)
		//}
		game.Tick()

		time.Sleep(duration)
	}
}

func printGame(game *GameInstance, i int) {
	fmt.Println(strings.Repeat("=", len(game.String())))
	fmt.Printf("Tick: %d\n", i)
	fmt.Println(game)
}
