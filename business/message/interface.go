package message

import "time"

type Service interface {
	FindMessageByCreatedDate(date time.Time) (*Message, error)
}

type Repository interface {
	FindMessageByCreatedDate(date time.Time) (*Message, error)
}
