package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Nastyyy/mdm-back/market"
	uuid "github.com/satori/go.uuid"
)

const DEBUG bool = true

// Generates random UUID and writes response with it
func authorize(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")

	cookie, err := r.Cookie("uuid")
	if err != nil {
		fmt.Println("User does not have uuid")
		uid := uuid.NewV4()
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
	} else {
		os.Setenv("DEBUG", "false")
	}

	tickRate := time.Minute * 2
	if DEBUG {
		tickRate = time.Second * 1
	}

	admin, _ := market.NewUser("admin", uuid.NewV4().String())

	gameSession := market.NewSession(admin)

	game := market.GameInstance{Running: true, ID: 1, TickRate: tickRate, Market: market.NewMarket()}
	gameSession.SetGameInstance(&game)

	// TODO: Possibly clean up and move to Game.go?
	go func() {
		fmt.Println("Starting market game...")
		for game.Running {
			game.Tick()
			gameSession.SyncState()
			if DEBUG {
				//fmt.Println(game)
			}
			time.Sleep(game.TickRate)
		}
	}()

	fmt.Println("Starting HTTP server...")
	http.HandleFunc("/", gameSession.SocketHandler)
	http.HandleFunc("/authorize", authorize)
	http.HandleFunc("/getusername", getUsername)
	s := &http.Server{
		Addr: ":8080",
	}

	log.Fatal(s.ListenAndServe())
}
