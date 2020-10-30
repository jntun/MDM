package market

import (
	"encoding/json"
	"fmt"
	"os"

	uuid "github.com/satori/go.uuid"
)

type Message struct {
	UUID   uuid.UUID `json:"uuid"`
	Action string    `json:"action"`

	// TODO: Body handler
	Body interface{} `json:"body"`
}

func NewMessage(byteMsg []byte) *Message {
	if os.Getenv("DEBUG") == "true" {
		fmt.Println(string(byteMsg))
	}

	msg := Message{}
	err := json.Unmarshal(byteMsg, &msg)
	if err != nil {
		fmt.Println("NewMessage() - Error unmarshaling message: ", err, "\nMessage: ", string(byteMsg))
	}

	return &msg
}
