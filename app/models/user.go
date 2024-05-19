package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key"`
	Username          string    `gorm:"uniqueIndex;not null"`
	Email             string    `gorm:"uniqueIndex;not null"`
	Password          string    `gorm:"not null"`
	IsPremium         bool
	PremiumExpiryTime time.Time
	gorm.Model
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}
