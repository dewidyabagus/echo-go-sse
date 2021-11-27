package message

import "time"

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) FindMessageByCreatedDate(date time.Time) (*Message, error) {
	return s.repository.FindMessageByCreatedDate(time.Now())
}
