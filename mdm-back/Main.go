package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Nastyyy/mdm-back/config"
	"github.com/Nastyyy/mdm-back/market"
	uuid "github.com/satori/go.uuid"
)

const DEBUG bool = true

// Generates random UUID and writes response with it
func authorize(w http.ResponseWriter, r *http.Request) {
    (w).Header().Set("Access-Control-Allow-Origin", "*")

    cookie, err := r.Cookie("uuid")
    if err != nil {
            uid := uuid.NewV4()
            log.Printf("User does not have uuid - assigning: %s", uid)
            w.Write([]byte(uid.String()))
    } else {
            fmt.Println(cookie)
    }
}

func getUsername(w http.ResponseWriter, r *http.Request) {
    (w).Header().Set("Access-Control-Allow-Origin", "*")

    cookie, err := r.Cookie("uuid")
    if err != nil {
            log.Printf("/getusername: Error getting uuid cookie %v", err)
    }

    // TODO: Search session for user with matching uuid and write username
    fmt.Println(cookie)
}

func main() {
    if DEBUG {
            os.Setenv("DEBUG", "true")
            config.EnableDebug()
            //config.DEBUG_PERF = true
            //config.EnableAllDebug()
    } 

    ticker := config.Ticker()

    /* Default admin and session */
    admin, _ := market.NewUser("admin", uuid.NewV4().String())
    gameSession := market.NewSession(admin)

    /* Game instance */
    game := market.GameInstance{Running: true, ID: 1, Ticker: ticker, Market: market.NewMarket()}
    gameSession.SetGameInstance(&game)

    /* Game instance loop */
    // TODO: Possibly clean up and move to Game.go?
    fmt.Println("[Main] Starting market game...")
    go func() {
        for range ticker.C {
            startGameTime := time.Now()
            if gameSession.Game.Running {
                gameSession.Game.Tick()
            } else {
                config.VerboseLog(fmt.Sprintf("[Game-%d] Skipping game tick while paused...", gameSession.Game.ID))
            }

            gameSession.SyncState()
            endGameTime := time.Now()
            finalTime := endGameTime.Sub(startGameTime)

            config.PerfLog(fmt.Sprintf("[Game-%d] Tick took %v", gameSession.Game.ID, finalTime))
        }
    }()

    // TODO: possible cmd interface?

    fmt.Println("[Main] Starting HTTP server...")
    http.HandleFunc("/", gameSession.SocketHandler)
    http.HandleFunc("/authorize", authorize)
    http.HandleFunc("/getusername", getUsername)

    s := &http.Server{
            Addr: ":8080",
    }

    fmt.Println("")
    log.Fatal(s.ListenAndServe())
}
