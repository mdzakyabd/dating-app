package usecase

import (
	"github.com/google/uuid"
	"github.com/mdzakyabd/dating-app/app/models"
	"github.com/mdzakyabd/dating-app/app/repository"
)

type SwipeUseCase interface {
	Swipe(swipe *models.Swipe) error
}

type swipeUseCase struct {
	swipeRepo repository.SwipeRepository
	matchRepo repository.MatchRepository
}

func NewSwipeUseCase(swipeRepo repository.SwipeRepository, matchRepo repository.MatchRepository) SwipeUseCase {
	return &swipeUseCase{swipeRepo, matchRepo}
}

func (uc *swipeUseCase) Swipe(swipe *models.Swipe) error {

	if err := uc.swipeRepo.CreateSwipe(swipe); err != nil {
		return err
	}

	if swipe.Liked {
		// Check if there is a mutual like
		target, err := uc.swipeRepo.GetSwipe(swipe.TargetUserID, swipe.UserID)
		if err == nil && target.Liked {
			matchID := uuid.New()
			match := &models.MatchRoom{
				ID:           matchID,
				UserID:       swipe.UserID,
				TargetUserID: swipe.TargetUserID,
			}
			uc.matchRepo.CreateMatch(match)

			match2 := &models.MatchRoom{
				ID:           matchID,
				UserID:       swipe.TargetUserID,
				TargetUserID: swipe.UserID,
			}
			uc.matchRepo.CreateMatch(match2)
		}
	}
	return nil
}
