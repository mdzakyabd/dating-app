package tests

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/mdzakyabd/dating-app/app/models"
	"github.com/mdzakyabd/dating-app/app/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockMatchRepository struct {
	mock.Mock
}

// CreateMatch is a mocked implementation of the CreateMatch method in the MatchRepository interface
func (m *MockMatchRepository) CreateMatch(match *models.MatchRoom) error {
	args := m.Called(match)
	return args.Error(0)
}

// GetMatchedUsersID is a mocked implementation of the GetMatchedUsersID method in the MatchRepository interface
func (m *MockMatchRepository) GetMatchedUsersID(userID uuid.UUID, date time.Time) ([]uuid.UUID, error) {
	args := m.Called(userID, date)
	return args.Get(0).([]uuid.UUID), args.Error(1)
}

// GetMatchRooms is a mocked implementation of the GetMatchRooms method in the MatchRepository interface
func (m *MockMatchRepository) GetMatchRooms(userID uuid.UUID) ([]models.MatchRoom, error) {
	args := m.Called(userID)
	return args.Get(0).([]models.MatchRoom), args.Error(1)
}

// DeleteMatchRoom is a mocked implementation of the DeleteMatchRoom method in the MatchRepository interface
func (m *MockMatchRepository) DeleteMatchRoom(id, userID uuid.UUID) error {
	args := m.Called(id, userID)
	return args.Error(0)
}

// CreateMessage is a mocked implementation of the CreateMessage method in the MatchRepository interface
func (m *MockMatchRepository) CreateMessage(message *models.Message) error {
	args := m.Called(message)
	return args.Error(0)
}

// GetMessages is a mocked implementation of the GetMessages method in the MatchRepository interface
func (m *MockMatchRepository) GetMessages(matchRoomID uuid.UUID) ([]models.Message, error) {
	args := m.Called(matchRoomID)
	return args.Get(0).([]models.Message), args.Error(1)
}

func TestGetMatchRooms(t *testing.T) {
	// Create a new instance of the mock MatchRepository
	mockMatchRepo := new(MockMatchRepository)
	matchUseCase := usecase.NewMatchUsecase(mockMatchRepo)

	// Mock user ID
	userID := uuid.New()

	// Mock match rooms
	mockMatchRooms := []models.MatchRoom{
		{ID: uuid.New(), UserID: userID},
		{ID: uuid.New(), UserID: userID},
		// Add more match rooms as needed
	}

	// Set up expectation for GetMatchRooms method in mock repository
	mockMatchRepo.On("GetMatchRooms", userID).Return(mockMatchRooms, nil)

	// Call the GetMatchRooms method and assert the result
	resultMatchRooms, err := matchUseCase.GetMatchRooms(userID)
	assert.NoError(t, err)
	assert.NotNil(t, resultMatchRooms)
	assert.Len(t, resultMatchRooms, len(mockMatchRooms))

	// Assert that all expectations were met
	mockMatchRepo.AssertExpectations(t)
}

func TestDeleteMatchRoom(t *testing.T) {
	// Create a new instance of the mock MatchRepository
	mockMatchRepo := new(MockMatchRepository)
	matchUseCase := usecase.NewMatchUsecase(mockMatchRepo)

	// Mock match room ID and user ID
	matchRoomID := uuid.New()
	userID := uuid.New()

	// Set up expectation for DeleteMatchRoom method in mock repository
	mockMatchRepo.On("DeleteMatchRoom", matchRoomID, userID).Return(nil)

	// Call the DeleteMatchRoom method and assert the result
	err := matchUseCase.DeleteMatchRoom(matchRoomID, userID)
	assert.NoError(t, err)

	// Assert that all expectations were met
	mockMatchRepo.AssertExpectations(t)
}

func TestCreateMessage(t *testing.T) {
	// Create a new instance of the mock MatchRepository
	mockMatchRepo := new(MockMatchRepository)
	matchUseCase := usecase.NewMatchUsecase(mockMatchRepo)

	// Mock message
	mockMessage := &models.Message{
		ID:          uuid.New(),
		MatchRoomID: uuid.New(),
		SenderID:    uuid.New(),
		Content:     "Hello, World!",
		// Add other message fields as needed
	}

	// Set up expectation for CreateMessage method in mock repository
	mockMatchRepo.On("CreateMessage", mockMessage).Return(nil)

	// Call the CreateMessage method and assert the result
	err := matchUseCase.CreateMessage(mockMessage)
	assert.NoError(t, err)

	// Assert that all expectations were met
	mockMatchRepo.AssertExpectations(t)
}

func TestGetMessages(t *testing.T) {
	// Create a new instance of the mock MatchRepository
	mockMatchRepo := new(MockMatchRepository)
	matchUseCase := usecase.NewMatchUsecase(mockMatchRepo)

	// Mock match room ID
	matchRoomID := uuid.New()

	// Mock messages
	mockMessages := []models.Message{
		{ID: uuid.New(), MatchRoomID: matchRoomID},
		{ID: uuid.New(), MatchRoomID: matchRoomID},
		// Add more messages as needed
	}

	// Set up expectation for GetMessages method in mock repository
	mockMatchRepo.On("GetMessages", matchRoomID).Return(mockMessages, nil)

	// Call the GetMessages method and assert the result
	resultMessages, err := matchUseCase.GetMessages(matchRoomID)
	assert.NoError(t, err)
	assert.NotNil(t, resultMessages)
	assert.Len(t, resultMessages, len(mockMessages))

	// Assert that all expectations were met
	mockMatchRepo.AssertExpectations(t)
}
