package market

import (
	"fmt"
	"log"
	"net/http"
	"strings"

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
	Admin *User            `json:"admin,omitempty"`
	Users map[string]*User `json:"users,omitempty"`
	Game  *GameInstance    `json:"game"`
}

func (sess *Session) SetGameInstance(game *GameInstance) {
	sess.Game = game
	fmt.Println(sess.Game)
}

func (sess *Session) AddUser(user *User) {
	//sess.Users = append(sess.Users, user)
	sess.Users[user.UUID.String()] = user
}

func (sess *Session) GetUser(uuid string) *User {
	return sess.Users[uuid]
}

func (sess *Session) SocketHandler(w http.ResponseWriter, r *http.Request) {
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

			if err = sess.handleMessage(conn, msg); err != nil {
				log.Println(err)
				break
			}
		}
	}()
}

// SyncState writes the session state to all given
// player connections
// sess.syncing is a posssible work-around for concurrent writes to connections
// not sure if actually works
func (sess *Session) SyncState() {
	//fmt.Println("Syncing state...")
	for _, user := range sess.Users {
		user.SendUpdate(sess)
	}
}

func (sess *Session) handleMessage(conn *websocket.Conn, byteMsg []byte) error {
	msg := NewMessage(byteMsg)
	fmt.Println(strings.Repeat("-", 100))

	MapEvent(sess, conn, msg)
	return nil
}

func NewSession(admin *User) *Session {
	session := Session{Admin: admin}
	users := make(map[string]*User)
	session.Users = users
	return &session
}

func (sess *Session) String() string {
	return fmt.Sprintf("Admin: %s | Players: %v\n%v", sess.Admin, sess.Users, sess.Game.Market)
}
