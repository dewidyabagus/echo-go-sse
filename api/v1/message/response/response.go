package response

import (
	"time"

	"go-sse/business/message"
)

type Message struct {
	Message   string
	CreatedAt time.Time
}

type Dashboard struct {
	Name   string
	Events chan *Message
}

func ResponseMessage(m *message.Message) *Message {
	return &Message{
		Message:   m.Message,
		CreatedAt: m.CreatedAt,
	}
}
