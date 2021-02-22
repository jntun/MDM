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
            config.DebugLog(fmt.Sprintf("[Main] User does not have uuid - assigning: %s", uid))
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
    args := os.Args[1:]
    if len(args) != 0 {
        flags(args)
    }

    /* Default admin and session */
    //admin, _ := market.NewUser("admin", uuid.NewV4().String())
    admin, _ := market.NewUser("admin", "a5894330-c1b0-4115-9a41-5897bd9a291c")
    gameSession := market.NewSession(admin)

    /* Game instance */
    game := market.GameInstance{Running: true, ID: 1, Ticker: config.Ticker(), Market: market.NewMarket()}
    gameSession.SetGameInstance(&game)

    /* Game instance loop */
    // TODO: Possibly clean up and move to Game.go?
    fmt.Println("[Main] Starting market game...")
    go func() {
        for range gameSession.Game.Ticker.C {
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

func flags(args []string) {
    config.FlagLog("[flags] Processing flags...")
    for _, arg := range args {
        config.FlagLog(arg)
        if arg[0] == byte('-') {
            switch string(arg[1]) {
            case "d":
                config.DEBUG = true
                config.DebugLog("Enabled log.")
            case "v":
                config.DEBUG_VERBOSE = true
                config.VerboseLog("Enabled log.")
            case "g":
                /* Game debug */

                // Stock debug
                config.DEBUG_STOCK = true
                config.StockLog("initialized in debug mode")
            case "p":
                config.DEBUG_PERF = true
                config.PerfLog("Enabled log.")
            }
        }
    }
}
