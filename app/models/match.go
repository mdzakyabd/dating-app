package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MatchRoom struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key"`
	UserID       uuid.UUID `gorm:"type:uuid;not null"`
	TargetUserID uuid.UUID `gorm:"type:uuid;not null"`
	gorm.Model
}
