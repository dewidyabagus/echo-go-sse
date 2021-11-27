package message

import "time"

type Message struct {
	ID        int
	CreatedAt time.Time
	Message   string
}
