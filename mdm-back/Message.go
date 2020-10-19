package main

import (
	"encoding/json"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type Message struct {
	UUID   uuid.UUID `json:"uuid"`
	Action string    `json:"action"`

	// TODO: Body handler
	Body interface{} `json:"body"`
}

func NewMessage(byteMsg []byte) *Message {
	if DEBUG {
		fmt.Println(string(byteMsg))
	}

	msg := Message{}
	err := json.Unmarshal(byteMsg, &msg)
	if err != nil {
		fmt.Println("Error unmarshaling message: ", err, "\nMessage: ", string(byteMsg))
	}

	return &msg
}
