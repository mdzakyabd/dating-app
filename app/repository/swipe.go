package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/mdzakyabd/dating-app/app/models"
	"gorm.io/gorm"
)

type SwipeRepository interface {
	GetSwipedUsersID(userID uuid.UUID, date time.Time) ([]uuid.UUID, error)
	CreateSwipe(swipe *models.Swipe) error
	GetSwipe(userID, targetUserID uuid.UUID) (*models.Swipe, error)
}

type swipeRepository struct {
	db *gorm.DB
}

func NewSwipeRepository(db *gorm.DB) SwipeRepository {
	return &swipeRepository{db: db}
}

func (r *swipeRepository) CreateSwipe(swipe *models.Swipe) error {
	return r.db.Create(swipe).Error
}

func (r *swipeRepository) GetSwipedUsersID(userID uuid.UUID, date time.Time) ([]uuid.UUID, error) {
	var swipes []models.Swipe
	err := r.db.Select("target_user_id").Where("user_id = ? AND DATE(created_at) = ?", userID, date).Find(&swipes).Error
	if err != nil {
		return nil, err
	}

	SwipedUserID := make([]uuid.UUID, len(swipes))
	for i, swipe := range swipes {
		SwipedUserID[i] = swipe.TargetUserID
	}

	return SwipedUserID, nil
}

func (r *swipeRepository) GetSwipe(userID, targetUserID uuid.UUID) (*models.Swipe, error) {
	var swipe models.Swipe
	err := r.db.Where("user_id = ? AND target_user_id = ?", userID, targetUserID).
		Order("created_at desc").
		First(&swipe).Error
	if err != nil {
		return nil, err
	}
	return &swipe, nil
}
