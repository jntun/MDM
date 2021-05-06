package market

import (
	"encoding/json"
	"fmt"

	"github.com/Nastyyy/mdm-back/config"
	uuid "github.com/satori/go.uuid"
)

type Message struct {
	UUID   uuid.UUID `json:"uuid"`
	Action string    `json:"action"`

	// TODO: Body handler
	Body interface{} `json:"body"`
}

func (msg *Message) String() string {
	return fmt.Sprintf("UUID: %s | Action: \"%s\" | Body: %s", msg.UUID, msg.Action, msg.Body)
}

func NewMessage(byteMsg []byte) *Message {
	config.VerboseLog(fmt.Sprintf("[NewMessage] %s", byteMsg))

	msg := Message{}
	err := json.Unmarshal(byteMsg, &msg)
	if err != nil {
		config.DebugLog(fmt.Sprintf("[NewMessage] - Error unmarshaling message: ", err, string(byteMsg)))
	}

	return &msg
}
