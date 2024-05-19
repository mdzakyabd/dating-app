package tests

import (
	"testing"

	"github.com/google/uuid"
	"github.com/mdzakyabd/dating-app/app/models"
	"github.com/mdzakyabd/dating-app/app/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking dependencies
type MockProfileRepository struct {
	mock.Mock
}

// CreateProfile is a mocked implementation of the CreateProfile method in the ProfileRepository interface
func (m *MockProfileRepository) CreateProfile(profile *models.Profile) error {
	args := m.Called(profile)
	return args.Error(0)
}

// GetProfileByID is a mocked implementation of the GetProfileByID method in the ProfileRepository interface
func (m *MockProfileRepository) GetProfileByID(id uuid.UUID) (*models.Profile, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Profile), args.Error(1)
}

// UpdateProfile is a mocked implementation of the UpdateProfile method in the ProfileRepository interface
func (m *MockProfileRepository) UpdateProfile(profile *models.Profile) error {
	args := m.Called(profile)
	return args.Error(0)
}

// GetProfilesExcluding is a mocked implementation of the GetProfilesExcluding method in the ProfileRepository interface
func (m *MockProfileRepository) GetProfilesExcluding(excludeIDs []uuid.UUID, limit int) ([]models.Profile, error) {
	args := m.Called(excludeIDs, limit)
	return args.Get(0).([]models.Profile), args.Error(1)
}

// Test cases
func TestValidProfileCreation(t *testing.T) {
	mockProfileRepo := new(MockProfileRepository)
	mockUserRepo := new(MockUserRepository)
	mockMatchRepo := new(MockMatchRepository)
	mockSwipeRepo := new(MockSwipeRepository)
	// Create mock repositories for dependencies

	// Create a new instance of the ProfileUseCase with the mock repositories
	profileUseCase := usecase.NewProfileUseCase(mockProfileRepo, mockUserRepo, mockSwipeRepo, mockMatchRepo)

	// Valid profile creation input
	profile := &models.Profile{
		ID:           uuid.New(),
		UserID:       uuid.New(),
		Name:         "Test User",
		Bio:          "This is a test bio",
		ProfileImage: "http://example.com/image.jpg",
	}

	mockProfileRepo.On("CreateProfile", profile).Return(nil)

	err := profileUseCase.CreateProfile(profile)

	assert.Nil(t, err)
	mockProfileRepo.AssertExpectations(t)
}

func TestValidProfileUpdate(t *testing.T) {
	mockProfileRepo := new(MockProfileRepository)
	mockUserRepo := new(MockUserRepository)
	mockMatchRepo := new(MockMatchRepository)
	mockSwipeRepo := new(MockSwipeRepository)
	// Create mock repositories for dependencies

	// Create a new instance of the ProfileUseCase with the mock repositories
	profileUseCase := usecase.NewProfileUseCase(mockProfileRepo, mockUserRepo, mockSwipeRepo, mockMatchRepo)

	// Valid profile update input
	profile := &models.Profile{
		ID:           uuid.New(),
		UserID:       uuid.New(),
		Name:         "Updated User",
		Bio:          "This is an updated bio",
		ProfileImage: "http://example.com/newimage.jpg",
	}

	mockProfileRepo.On("UpdateProfile", profile).Return(nil)

	err := profileUseCase.UpdateProfile(profile)

	assert.Nil(t, err)
	mockProfileRepo.AssertExpectations(t)
}

func TestGetProfileByID(t *testing.T) {
	mockProfileRepo := new(MockProfileRepository)
	mockUserRepo := new(MockUserRepository)
	mockMatchRepo := new(MockMatchRepository)
	mockSwipeRepo := new(MockSwipeRepository)
	// Create mock repositories for dependencies

	// Create a new instance of the ProfileUseCase with the mock repositories
	profileUseCase := usecase.NewProfileUseCase(mockProfileRepo, mockUserRepo, mockSwipeRepo, mockMatchRepo)

	// Mock a profile
	profileID := uuid.New()
	mockProfile := &models.Profile{
		ID:     profileID,
		Name:   "Test Profile",
		UserID: uuid.New(),
		// Add other profile fields as needed
	}

	// Set up expectation for GetProfileByID method in mock repository
	mockProfileRepo.On("GetProfileByID", profileID).Return(mockProfile, nil)

	// Call the GetProfileByID method and assert the result
	resultProfile, err := profileUseCase.GetProfileByID(profileID)
	assert.NoError(t, err)
	assert.NotNil(t, resultProfile)
	assert.Equal(t, profileID, resultProfile.ID)

	// Assert that all expectations were met
	mockProfileRepo.AssertExpectations(t)
}

func TestViewProfiles(t *testing.T) {
	mockProfileRepo := new(MockProfileRepository)
	mockUserRepo := new(MockUserRepository)
	mockMatchRepo := new(MockMatchRepository)
	mockSwipeRepo := new(MockSwipeRepository)
	// Create mock repositories for dependencies

	// Create a new instance of the ProfileUseCase with the mock repositories
	profileUseCase := usecase.NewProfileUseCase(mockProfileRepo, mockUserRepo, mockSwipeRepo, mockMatchRepo)

	// Mock user ID
	userID := uuid.New()

	// Mock user with premium status
	mockUser := &models.User{
		ID:        userID,
		IsPremium: true,
	}

	// Set up expectations for GetUserByID method in mock user repository
	mockUserRepo.On("GetUserByID", userID).Return(mockUser, nil)

	// Set up expectations for GetSwipedUsersID method in mock swipe repository
	mockSwipeRepo.On("GetSwipedUsersID", userID, mock.Anything).Return([]uuid.UUID{}, nil)

	// Set up expectations for GetMatchedUsersID method in mock match repository
	mockMatchRepo.On("GetMatchedUsersID", userID, mock.Anything).Return([]uuid.UUID{}, nil)

	// Mock profiles
	mockProfiles := []models.Profile{
		{ID: uuid.New(), Name: "Profile 1", UserID: uuid.New()},
		{ID: uuid.New(), Name: "Profile 2", UserID: uuid.New()},
		// Add more profiles as needed
	}

	// Set up expectations for GetProfilesExcluding method in mock profile repository
	mockProfileRepo.On("GetProfilesExcluding", mock.Anything, 50).Return(mockProfiles, nil)

	// Call the ViewProfiles method and assert the result
	resultProfiles, err := profileUseCase.ViewProfiles(userID)
	assert.NoError(t, err)
	assert.NotNil(t, resultProfiles)
	assert.Len(t, resultProfiles, len(mockProfiles))

	// Assert that all expectations were met
	mockUserRepo.AssertExpectations(t)
	mockSwipeRepo.AssertExpectations(t)
	mockMatchRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
}
