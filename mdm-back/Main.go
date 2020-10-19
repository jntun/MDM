package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Nastyyy/mdm-back/market"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

const DEBUG bool = true

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func session(w http.ResponseWriter, r *http.Request) {
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
		handleMessage(msg)
	}
}

func handleMessage(byteMsg []byte) {
	msg := NewMessage(byteMsg)
	EventHandler(msg)
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

func main() {
	if DEBUG {
		os.Setenv("DEBUG", "true")
	}

	market.Game()

	http.HandleFunc("/", session)
	http.HandleFunc("/authorize", authorize)
	s := &http.Server{
		Addr: ":8080",
	}

	fmt.Println("Serve")
	log.Fatal(s.ListenAndServe())

}
