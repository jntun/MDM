package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

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
		}
		handleMessage(msg)
	}
}

func handleMessage(msg []byte) {
	fmt.Println(msg)
	/*
		var data map[string]interface{}
		err := json.Unmarshal(msg, &data)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(data["uuid"].(string), data["body"])
		switch data["body"] {
		case "register":
			fmt.Println("Registering")
		}
	*/
}

func newCookie() *http.Cookie {
	uid := uuid.NewV4()
	return &http.Cookie{
		Name:    "uuid",
		Value:   uid.String(),
		Expires: time.Now().Add(365 * 24 * time.Hour),
	}
}

func authorize(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")

	cookie, err := r.Cookie("uuid")
	if err != nil {
		fmt.Println("User does not have uuid")
		http.SetCookie(w, newCookie())
	} else {
		fmt.Println(cookie)
	}
}

func main() {
	http.HandleFunc("/", session)
	http.HandleFunc("/authorize", authorize)
	s := &http.Server{
		Addr: ":8080",
	}
	log.Fatal(s.ListenAndServe())
}
