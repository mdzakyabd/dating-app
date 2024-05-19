package usecase

import (
	"github.com/google/uuid"
	"github.com/mdzakyabd/dating-app/app/models"
	"github.com/mdzakyabd/dating-app/app/repository"
)

type MatchUsecase interface {
	GetMatchRooms(userID uuid.UUID) ([]models.MatchRoom, error)
	DeleteMatchRoom(id, userID uuid.UUID) error
	CreateMessage(matchRoom *models.Message) error
	GetMessages(matchRoomID uuid.UUID) ([]models.Message, error)
}

type matchUsecase struct {
	matchRepo repository.MatchRepository
}

func NewMatchUsecase(repo repository.MatchRepository) MatchUsecase {
	return &matchUsecase{repo}
}

func (u *matchUsecase) GetMatchRooms(userID uuid.UUID) ([]models.MatchRoom, error) {
	return u.matchRepo.GetMatchRooms(userID)
}

func (u *matchUsecase) DeleteMatchRoom(id, userID uuid.UUID) error {
	return u.matchRepo.DeleteMatchRoom(id, userID)
}

func (u *matchUsecase) CreateMessage(matchRoom *models.Message) error {
	return u.matchRepo.CreateMessage(matchRoom)
}

func (u *matchUsecase) GetMessages(matchRoomID uuid.UUID) ([]models.Message, error) {
	return u.matchRepo.GetMessages(matchRoomID)
}
