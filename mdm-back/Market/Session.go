package market

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Session represents the logic that connects all users for a given instance
type Session struct {
	Admin *User         `json:"admin,omitempty"`
	Users []*User       `json:"users,omitempty"`
	Game  *GameInstance `json:"game"`
}

func (sess *Session) SetGameInstance(game *GameInstance) {
	sess.Game = game
	fmt.Println(sess.Game)
}

func (sess *Session) AddUser(user *User) {
	sess.Users = append(sess.Users, user)
}

func (sess *Session) GetUser(cookieUUID string) (user *User, err error) {
	for _, user := range sess.Users {
		if user.UUID.String() == cookieUUID {
			return user, nil
		}
	}
	return nil, errors.New("Could not match player with UUID " + cookieUUID)
}

func (sess Session) SocketHandler(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				break
			}

			if err = sess.handleMessage(msg); err != nil {
				log.Println(err)
				break
			}
		}
	}()

	go func() {
		for true {
			fmt.Println("Sending update...")
			conn.WriteJSON(sess)
			time.Sleep(time.Second * 1)
		}
	}()
}

// SyncState writes the session state to all given
// player connections
func (sess Session) SyncState() {
	for _, user := range sess.Users {
		if user.Conn != nil {
			user.Conn.WriteJSON(sess)
		} else {
			log.Printf("%s does not have active connection object. Unable to sync", user.Name)
		}
	}
}

func (sess Session) handleMessage(byteMsg []byte) error {
	msg := NewMessage(byteMsg)
	MapEvent(&sess, msg)
	fmt.Println(strings.Repeat("-", 80))
	return nil
}

func NewSession(admin *User) *Session {
	session := Session{Admin: admin}
	return &session
}

func (sess Session) String() string {
	return fmt.Sprintf("Admin: %s | Players: %v\n%v", sess.Admin, sess.Users, sess.Game.Market)
}
