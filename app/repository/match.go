package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/mdzakyabd/dating-app/app/models"
	"gorm.io/gorm"
)

type MatchRepository interface {
	CreateMatch(match *models.MatchRoom) error
	GetMatchedUsersID(userID uuid.UUID, date time.Time) ([]uuid.UUID, error)
	GetMatchRooms(userID uuid.UUID) ([]models.MatchRoom, error)
	DeleteMatchRoom(id, userID uuid.UUID) error
	CreateMessage(message *models.Message) error
	GetMessages(matchRoomID uuid.UUID) ([]models.Message, error)
}

type matchRepository struct {
	db *gorm.DB
}

func NewMatchRepository(db *gorm.DB) MatchRepository {
	return &matchRepository{db: db}
}

func (r *matchRepository) CreateMatch(match *models.MatchRoom) error {
	return r.db.Create(match).Error
}

func (r *matchRepository) GetMatchedUsersID(userID uuid.UUID, date time.Time) ([]uuid.UUID, error) {
	var matches []models.MatchRoom
	err := r.db.Select("target_user_id").Where("user_id = ? AND DATE(created_at) = ?", userID, date).Find(&matches).Error
	if err != nil {
		return nil, err
	}

	MatchedUserID := make([]uuid.UUID, len(matches))
	for i, match := range matches {
		MatchedUserID[i] = match.TargetUserID
	}

	return MatchedUserID, nil
}

func (r *matchRepository) GetMatchRooms(userID uuid.UUID) ([]models.MatchRoom, error) {
	var matchRooms []models.MatchRoom
	err := r.db.Where("user_id = ?", userID).Find(&matchRooms).Error
	return matchRooms, err
}

func (r *matchRepository) DeleteMatchRoom(id, userID uuid.UUID) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.MatchRoom{}).Error
}

func (r *matchRepository) CreateMessage(message *models.Message) error {
	return r.db.Create(message).Error
}

func (r *matchRepository) GetMessages(matchRoomID uuid.UUID) ([]models.Message, error) {
	var messages []models.Message
	err := r.db.Where("match_room_id = ?", matchRoomID).Find(&messages).Error
	return messages, err
}
