package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	MatchRoomID uuid.UUID `gorm:"type:uuid;not null"`
	SenderID    uuid.UUID `gorm:"type:uuid;not null"`
	Content     string    `gorm:"not null"`
	gorm.Model
}

func (message *Message) BeforeCreate(tx *gorm.DB) (err error) {
	message.ID = uuid.New()
	return
}
