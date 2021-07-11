package stark

import (
	"fmt"
	"log"

	"github.com/Nastyyy/mdm-back/market"
	"github.com/gorilla/websocket"
	suuid "github.com/satori/go.uuid"
)

type Client struct {
	uuid *suuid.UUID
	conn *websocket.Conn
}

func (sc *Client) start() {
	err := sc.sendEvent(market.REGISTER, &market.RegisterAction{UUID: sc.uuid.String()})
	if err != nil {
		StarkError("sendEvent(): %s", err)
	}

	type t interface{}
	body := map[string]t{
		"test": true,
	}
	err = sc.sendEvent(market.ADMIN, &market.AdminAction{Body: body})
	if err != nil {
		StarkError("sendEvent(): %s", err)
	}

	for {
		_, message, err := sc.conn.ReadMessage()
		if err != nil {
			StarkError("ReadMessage(): %s", err)
		}
		StarkLog("[Message]: %v", string(message))
	}
	return
}

func (sc *Client) sendEvent(action string, actCall market.Action) error {
	msg := &market.Message{*sc.uuid, action, actCall}
	StarkLog("sendEvent(): %s", msg)
	err := sc.conn.WriteJSON(msg)
	if err != nil {
		return err
	}

	return nil
}

// Gotcha
// If we return/treat we die...
func RunClient() int {
	var endpoint string = "ws://jntun.com:8080"
	c, _, err := websocket.DefaultDialer.Dial(endpoint, nil)
	if err != nil {
		StarkError("RunClient(): %s", err)
		return 1
	}

	defer c.Close()

	StarkLog("[CONN] Connected successfully to %s", endpoint)

	// Isn't it beautifully ugly in a way?
	// client := Client{&[]suuid.UUID{suuid.NewV4()}[0], c}

	/*              Temp hard coding for admin                      */
	uuid, _ := suuid.FromString("857bb89c-a8bf-4a64-92f6-c224307a4286")
	client := Client{&uuid, c}
	/*                                                              */

	client.start()
	return 0
}

func StarkLog(fstring string, vargs ...interface{}) {
	log.Printf("[STARK]%s\n", fmt.Sprintf(fstring, vargs))
}

func StarkError(fstring string, vargs ...interface{}) {
	log.Printf("[ERROR]%s\n", fmt.Sprintf(fstring, vargs))
}
