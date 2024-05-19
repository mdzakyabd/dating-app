package repository

import (
	"github.com/google/uuid"
	"github.com/mdzakyabd/dating-app/app/models"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	CreateProfile(profile *models.Profile) error
	GetProfileByID(id uuid.UUID) (*models.Profile, error)
	UpdateProfile(profile *models.Profile) error
	GetProfilesExcluding(excludeIDs []uuid.UUID, limit int) ([]models.Profile, error)
}
type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{db: db}
}

func (r *profileRepository) CreateProfile(profile *models.Profile) error {
	return r.db.Create(profile).Error
}

func (r *profileRepository) GetProfileByID(id uuid.UUID) (*models.Profile, error) {
	var profile models.Profile
	err := r.db.First(&profile, "id = ?", id).Error
	return &profile, err
}

func (r *profileRepository) UpdateProfile(profile *models.Profile) error {
	return r.db.Save(profile).Error
}

func (r *profileRepository) GetProfilesExcluding(excludeIDs []uuid.UUID, limit int) ([]models.Profile, error) {
	var profiles []models.Profile
	query := r.db.Where("user_id NOT IN ?", excludeIDs)

	err := query.Limit(limit).Find(&profiles).Error
	return profiles, err
}
