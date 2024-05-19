package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Swipe struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key"`
	UserID       uuid.UUID `gorm:"type:uuid;not null"`
	TargetUserID uuid.UUID `gorm:"type:uuid;not null"`
	Liked        bool      `gorm:"not null"`
	gorm.Model
}

func (swipe *Swipe) BeforeCreate(tx *gorm.DB) (err error) {
	swipe.ID = uuid.New()
	return
}
