package tests

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/mdzakyabd/dating-app/app/models"
	"github.com/mdzakyabd/dating-app/app/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking dependencies
type MockSwipeRepository struct {
	mock.Mock
}

// CreateSwipe is a mocked implementation of the CreateSwipe method in the SwipeRepository interface
func (m *MockSwipeRepository) CreateSwipe(swipe *models.Swipe) error {
	args := m.Called(swipe)
	return args.Error(0)
}

// GetSwipedUsersID is a mocked implementation of the GetSwipedUsersID method in the SwipeRepository interface
func (m *MockSwipeRepository) GetSwipedUsersID(userID uuid.UUID, date time.Time) ([]uuid.UUID, error) {
	args := m.Called(userID, date)
	return args.Get(0).([]uuid.UUID), args.Error(1)
}

// GetSwipe is a mocked implementation of the GetSwipe method in the SwipeRepository interface
func (m *MockSwipeRepository) GetSwipe(userID, targetUserID uuid.UUID) (*models.Swipe, error) {
	args := m.Called(userID, targetUserID)
	return args.Get(0).(*models.Swipe), args.Error(1)
}

func TestSwipe(t *testing.T) {
	// Create mock repositories
	mockMatchRepo := new(MockMatchRepository)
	mockSwipeRepo := new(MockSwipeRepository)

	// Create SwipeUseCase with mock repositories
	swipeUseCase := usecase.NewSwipeUseCase(mockSwipeRepo, mockMatchRepo)

	// Create a mock swipe
	swipe := &models.Swipe{
		ID:           uuid.New(),
		UserID:       uuid.New(),
		TargetUserID: uuid.New(),
		Liked:        true, // Simulating a swipe right (like)
	}

	// Set up expectation for CreateSwipe method in mock swipe repository
	mockSwipeRepo.On("CreateSwipe", swipe).Return(nil)

	// Set up expectation for GetSwipe method in mock swipe repository
	mockSwipeRepo.On("GetSwipe", swipe.TargetUserID, swipe.UserID).Return(&models.Swipe{}, nil)

	// Call the Swipe method and assert the result
	err := swipeUseCase.Swipe(swipe)
	assert.NoError(t, err)

	// Assert that all expectations were met
	mockSwipeRepo.AssertExpectations(t)
}

func TestSwipe_ErrorCreatingSwipe(t *testing.T) {
	// Create mock repositories
	mockMatchRepo := new(MockMatchRepository)
	mockSwipeRepo := new(MockSwipeRepository)

	// Create SwipeUseCase with mock repositories
	swipeUseCase := usecase.NewSwipeUseCase(mockSwipeRepo, mockMatchRepo)

	// Create a mock swipe
	swipe := &models.Swipe{
		ID:           uuid.New(),
		UserID:       uuid.New(),
		TargetUserID: uuid.New(),
		Liked:        true, // Simulating a swipe right (like)
	}

	// Set up expectation for CreateSwipe method in mock swipe repository to return an error
	mockSwipeRepo.On("CreateSwipe", swipe).Return(errors.New("error creating swipe"))

	// Call the Swipe method and assert the error
	err := swipeUseCase.Swipe(swipe)
	assert.Error(t, err)
	assert.EqualError(t, err, "error creating swipe")

	// Assert that all expectations were met
	mockSwipeRepo.AssertExpectations(t)
	mockMatchRepo.AssertExpectations(t)
}

func TestSwipe_NoMutualLike(t *testing.T) {
	// Create mock repositories
	mockMatchRepo := new(MockMatchRepository)
	mockSwipeRepo := new(MockSwipeRepository)

	// Create SwipeUseCase with mock repositories
	swipeUseCase := usecase.NewSwipeUseCase(mockSwipeRepo, mockMatchRepo)

	// Create a mock swipe
	swipe := &models.Swipe{
		ID:           uuid.New(),
		UserID:       uuid.New(),
		TargetUserID: uuid.New(),
		Liked:        true, // Simulating a swipe right (like)
	}

	// Set up expectation for CreateSwipe method in mock swipe repository
	mockSwipeRepo.On("CreateSwipe", swipe).Return(nil)

	// Set up expectation for GetSwipe method in mock swipe repository to return an error (no mutual like)
	mockSwipeRepo.On("GetSwipe", swipe.TargetUserID, swipe.UserID).Return(&models.Swipe{}, errors.New("no mutual like"))

	// Call the Swipe method and assert no error
	err := swipeUseCase.Swipe(swipe)
	assert.NoError(t, err)

	// Assert that all expectations were met
	mockSwipeRepo.AssertExpectations(t)
	mockMatchRepo.AssertExpectations(t)
}
