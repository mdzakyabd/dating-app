package usecase

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mdzakyabd/dating-app/app/models"
	"github.com/mdzakyabd/dating-app/app/repository"
	"github.com/mdzakyabd/dating-app/app/utils"
)

type UserUseCase interface {
	Register(user *models.User) error
	Login(email, password string) (*models.User, error)
	UpdateUser(user *models.User) error
	GetUserByID(id uuid.UUID) (*models.User, error)
	SubscribePremium(userID uuid.UUID, expiry time.Time) error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{userRepository: userRepository}
}

func (uc *userUseCase) Register(data *models.User) error {
	user, err := uc.userRepository.GetUserByEmail(data.Email)
	if err != nil {
		return err
	}

	if len(user.Email) > 0 {
		return errors.New("email already exists")
	}

	hashedPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return uc.userRepository.CreateUser(data)
}

func (uc *userUseCase) Login(email, password string) (*models.User, error) {
	user, err := uc.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

func (uc *userUseCase) UpdateUser(user *models.User) error {
	return uc.userRepository.UpdateUser(user)
}

func (uc *userUseCase) GetUserByID(id uuid.UUID) (*models.User, error) {
	return uc.userRepository.GetUserByID(id)
}

func (uc *userUseCase) SubscribePremium(userID uuid.UUID, expiry time.Time) error {
	return uc.userRepository.UpdatePremiumStatus(userID, true, expiry)
}
