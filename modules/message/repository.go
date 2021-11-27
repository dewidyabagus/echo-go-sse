package message

import (
	"time"

	"gorm.io/gorm"

	"go-sse/business/message"
)

type Message struct {
	ID        int       `gorm:"id;primaryKey;autoIncrement"`
	CreatedAt time.Time `gorm:"created_at;timestamp;index:posts_created_at_idx;not null"`
	Message   string    `gorm:"message;type:varchar(100)"`
}

type Repository struct {
	DB *gorm.DB
}

func (m *Message) toBusinessMessage() *message.Message {
	return &message.Message{
		ID:        m.ID,
		CreatedAt: m.CreatedAt,
		Message:   m.Message,
	}
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) FindMessageByCreatedDate(date time.Time) (*message.Message, error) {
	message := new(Message)

	if err := r.DB.First(message, "cast(created_at as date) = cast(? as date)", date).Error; err != nil {
		return nil, err
	}

	return message.toBusinessMessage(), nil
}
