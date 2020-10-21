package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Nastyyy/mdm-back/market"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

const DEBUG bool = false

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func session(w http.ResponseWriter, r *http.Request, sess *market.Session) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		handleMessage(msg, sess)
	}
}

func handleMessage(byteMsg []byte, sess *market.Session) {
	msg := NewMessage(byteMsg)
	EventHandler(msg, sess)
	fmt.Println(strings.Repeat("-", 80))
}

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
	}

	tickRate := time.Minute * 2
	if DEBUG {
		tickRate = time.Millisecond * 500
	}

	testUser := market.NewUser("admin")
	gameSession := market.NewSession(testUser)
	game := market.GameInstance{Running: true, TickRate: tickRate, Market: market.NewMarket()}
	gameSession.SetGameInstance(&game)

	go func() {
		fmt.Println("Starting market game...")
		for game.Running {
			game.Tick()
			if DEBUG {
				fmt.Println(game)
			}
			time.Sleep(game.TickRate)
		}
	}()

	fmt.Println("Starting HTTP server...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		session(w, r, gameSession)
	})
	http.HandleFunc("/authorize", authorize)
	http.HandleFunc("/getusername", getUsername)
	s := &http.Server{
		Addr: ":8080",
	}

	log.Fatal(s.ListenAndServe())
}
