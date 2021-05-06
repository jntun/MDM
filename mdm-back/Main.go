package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
        "strings"

	"github.com/Nastyyy/mdm-back/config"
	"github.com/Nastyyy/mdm-back/market"
        "github.com/Nastyyy/mdm-back/stark"
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
        admin, _ := market.NewUser("admin", "857bb89c-a8bf-4a64-92f6-c224307a4286")
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

                config.VerboseLog(fmt.Sprintf("Tick: %d", gameSession.Game.TickTotal))
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
            err := flagMatch(string(arg[1]))
            if err != nil {
                fmt.Printf("[Error] %s\n", err)
            }
        }
    }
}

func flagMatch(flag string) error {
        switch string(flag) {
            // C for client or simply a hacky debug helper which mimics a user
            case "C":
            /******* TAKES OVER CONTROL FLOW - NO RETURN IF ENABLED *******/
                starkUp()
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
            default:
                return fmt.Errorf("invalid flag provided -%s", flag)
        }
        return nil
}

// Headless, will never return
func starkUp() {
        big := strings.Repeat("=", 30)
        small := strings.Repeat("*", 10)
        fmt.Printf("%s\n%s Stark %s\n%s\n\n", big, small, small, big)

        ret := stark.RunClient()
        os.Exit(ret)
}
