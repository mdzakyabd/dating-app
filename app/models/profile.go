package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Profile struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key"`
	UserID       uuid.UUID `gorm:"type:uuid;not null"`
	Name         string
	Bio          string
	ProfileImage string
	gorm.Model
}

func (profile *Profile) BeforeCreate(tx *gorm.DB) (err error) {
	profile.ID = uuid.New()
	return
}
