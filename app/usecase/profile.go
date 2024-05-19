package usecase

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mdzakyabd/dating-app/app/models"
	"github.com/mdzakyabd/dating-app/app/repository"
	"github.com/mdzakyabd/dating-app/app/utils"
)

type ProfileUseCase interface {
	CreateProfile(profile *models.Profile) error
	GetProfileByID(id uuid.UUID) (*models.Profile, error)
	UpdateProfile(profile *models.Profile) error
	ViewProfiles(userID uuid.UUID) ([]models.Profile, error)
}

type profileUseCase struct {
	profileRepo repository.ProfileRepository
	userRepo    repository.UserRepository
	swipeRepo   repository.SwipeRepository
	matchRepo   repository.MatchRepository
}

func NewProfileUseCase(profileRepo repository.ProfileRepository, userRepo repository.UserRepository, swipeRepo repository.SwipeRepository, matchRepo repository.MatchRepository) ProfileUseCase {
	return &profileUseCase{profileRepo, userRepo, swipeRepo, matchRepo}
}

func (uc *profileUseCase) CreateProfile(profile *models.Profile) error {
	return uc.profileRepo.CreateProfile(profile)
}

func (uc *profileUseCase) GetProfileByID(id uuid.UUID) (*models.Profile, error) {
	return uc.profileRepo.GetProfileByID(id)
}

func (uc *profileUseCase) UpdateProfile(profile *models.Profile) error {
	return uc.profileRepo.UpdateProfile(profile)
}

func (uc *profileUseCase) ViewProfiles(userID uuid.UUID) ([]models.Profile, error) {
	user, err := uc.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	excludedUserID, err := uc.swipeRepo.GetSwipedUsersID(userID, time.Now())
	if err != nil {
		return nil, err
	}

	if !user.IsPremium && len(excludedUserID) >= 10 {
		return nil, errors.New("daily profile view limit reached")
	}

	matchedProfileIDs, err := uc.matchRepo.GetMatchedUsersID(userID, time.Now())
	if err != nil {
		return nil, err
	}

	excludedUserID = append(excludedUserID, matchedProfileIDs...)
	excludedUserID = append(excludedUserID, userID)

	excludedUserID = utils.DistinctUUIDs(excludedUserID)

	limit := 10
	if user.IsPremium {
		limit = 50 // Set a high limit for premium users to represent unlimited
	}

	profiles, err := uc.profileRepo.GetProfilesExcluding(excludedUserID, limit)
	if err != nil {
		return nil, err
	}

	return profiles, nil
}
