package tests

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/mdzakyabd/dating-app/app/models"
	"github.com/mdzakyabd/dating-app/app/usecase"
	"github.com/mdzakyabd/dating-app/app/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking dependencies
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByUsername(username string) (*models.User, error) {
	args := m.Called(username)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByID(id uuid.UUID) (*models.User, error) {
	args := m.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) UpdatePremiumStatus(userID uuid.UUID, isPremium bool, expiry time.Time) error {
	args := m.Called(userID, isPremium, expiry)
	return args.Error(0)
}

func (m *MockUserRepository) FindAllPremiumUsers() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

type MockUtils struct {
	mock.Mock
}

// CheckPasswordHash is a mocked implementation of the CheckPasswordHash function in the utils package
func (m *MockUtils) CheckPasswordHash(password, hashedPassword string) bool {
	args := m.Called(password, hashedPassword)
	return args.Bool(0)
}

func TestValidUserSignUp(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUsecase := usecase.NewUserUseCase(mockRepo)

	// Valid user input
	user := &models.User{
		ID:       uuid.New(),
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "securepassword",
	}

	user.Password, _ = utils.HashPassword(user.Password)

	mockRepo.On("GetUserByEmail", user.Email).Return(&models.User{}, nil)
	mockRepo.On("CreateUser", user).Return(nil)

	err := userUsecase.Register(user)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDuplicateEmailSignUp(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUsecase := usecase.NewUserUseCase(mockRepo)

	// Duplicate email input
	user := &models.User{
		ID:       uuid.New(),
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "securepassword",
	}

	mockRepo.On("GetUserByEmail", user.Email).Return(user, nil)

	err := userUsecase.Register(user)

	assert.NotNil(t, err)
	assert.Equal(t, "email already exists", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestLogin_ValidCredentials(t *testing.T) {
	// Create mock repository
	mockUserRepo := new(MockUserRepository)

	// Create UserUseCase with mock repository
	userUseCase := usecase.NewUserUseCase(mockUserRepo)

	// Define test data
	email := "test@example.com"
	password := "password"
	hashedPassword, _ := utils.HashPassword(password)
	user := &models.User{
		ID:       uuid.New(),
		Username: "testuser",
		Email:    email,
		Password: hashedPassword,
	}

	// Set up expectation for GetUserByEmail method in mock repository
	mockUserRepo.On("GetUserByEmail", email).Return(user, nil)

	// Call the Login method and assert the result
	resultUser, err := userUseCase.Login(email, password)
	assert.NoError(t, err)
	assert.NotNil(t, resultUser)
	assert.Equal(t, user, resultUser)

	// Assert that all expectations were met
	mockUserRepo.AssertExpectations(t)
}

func TestLogin_InvalidCredentials(t *testing.T) {
	// Create mock repository
	mockUserRepo := new(MockUserRepository)

	// Create UserUseCase with mock repository
	userUseCase := usecase.NewUserUseCase(mockUserRepo)

	// Define test data
	email := "test@example.com"
	password := "password"
	hashedPassword, _ := utils.HashPassword(password)
	user := &models.User{
		ID:       uuid.New(),
		Username: "testuser",
		Email:    email,
		Password: hashedPassword,
	}

	// Set up expectation for GetUserByEmail method in mock repository
	mockUserRepo.On("GetUserByEmail", email).Return(user, nil)

	// Call the Login method with invalid password and assert the error
	resultUser, err := userUseCase.Login(email, "invalidpassword")
	assert.Error(t, err)
	assert.Nil(t, resultUser)
	assert.EqualError(t, err, "invalid credentials")

	// Assert that all expectations were met
	mockUserRepo.AssertExpectations(t)
}

func TestLogin_UserNotFound(t *testing.T) {
	// Create mock repository
	mockUserRepo := new(MockUserRepository)

	// Create UserUseCase with mock repository
	userUseCase := usecase.NewUserUseCase(mockUserRepo)

	// Define test data
	email := "test@example.com"

	// Set up expectation for GetUserByEmail method in mock repository to return user not found error
	mockUserRepo.On("GetUserByEmail", email).Return(&models.User{}, errors.New("user not found"))

	// Call the Login method with non-existing email and assert the error
	resultUser, err := userUseCase.Login(email, "password")
	assert.Error(t, err)
	assert.Nil(t, resultUser)
	assert.EqualError(t, err, "user not found")

	// Assert that all expectations were met
	mockUserRepo.AssertExpectations(t)
}

func TestValidUpdateUserInfo(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUsecase := usecase.NewUserUseCase(mockRepo)

	// Valid update input
	user := &models.User{
		ID:       uuid.New(),
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "securepassword",
	}
	updatedInfo := &models.User{
		ID:       user.ID,
		Username: "newusername",
		Email:    "newemail@example.com",
		Password: "newsecurepassword",
	}

	mockRepo.On("UpdateUser", updatedInfo).Return(nil)

	err := userUsecase.UpdateUser(updatedInfo)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetUserByID(t *testing.T) {
	// Create mock repository
	mockUserRepo := new(MockUserRepository)

	// Create UserUseCase with mock repository
	userUseCase := usecase.NewUserUseCase(mockUserRepo)

	// Define test data
	userID := uuid.New()
	user := &models.User{
		ID:       userID,
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password",
	}

	// Set up expectation for GetUserByID method in mock repository
	mockUserRepo.On("GetUserByID", userID).Return(user, nil)

	// Call the GetUserByID method and assert the result
	resultUser, err := userUseCase.GetUserByID(userID)
	assert.NoError(t, err)
	assert.NotNil(t, resultUser)
	assert.Equal(t, user, resultUser)

	// Assert that all expectations were met
	mockUserRepo.AssertExpectations(t)
}

func TestSubscribePremium(t *testing.T) {
	// Create mock repository
	mockUserRepo := new(MockUserRepository)

	// Create UserUseCase with mock repository
	userUseCase := usecase.NewUserUseCase(mockUserRepo)

	// Define test data
	userID := uuid.New()
	expiry := time.Now().Add(24 * time.Hour) // Example expiry time

	// Set up expectation for UpdatePremiumStatus method in mock repository
	mockUserRepo.On("UpdatePremiumStatus", userID, true, expiry).Return(nil)

	// Call the SubscribePremium method and assert the result
	err := userUseCase.SubscribePremium(userID, expiry)
	assert.NoError(t, err)

	// Assert that all expectations were met
	mockUserRepo.AssertExpectations(t)
}
