package services

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"university_system/internal/domain/models"
)

type mockUserRepo struct {
	mock.Mock
}

func (m *mockUserRepo) GetUsers(ctx context.Context) ([]models.User, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *mockUserRepo) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *mockUserRepo) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	args := m.Called(ctx, username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *mockUserRepo) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	args := m.Called(ctx, user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *mockUserRepo) UpdateUser(ctx context.Context, user models.User) (*models.User, error) {
	args := m.Called(ctx, user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *mockUserRepo) DeleteUser(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestUserService_GetUsers(t *testing.T) {
	// Arrange
	mockRepo := new(mockUserRepo)
	svc := NewUserService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		expectedUsers := []models.User{
			{ID: "1", Username: "user1", Email: "user1@example.com"},
			{ID: "2", Username: "user2", Email: "user2@example.com"},
		}
		mockRepo.On("GetUsers", ctx).Return(expectedUsers, nil).Once()

		// Act
		users, err := svc.GetUsers(ctx)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedUsers, users)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("database error")
		mockRepo.On("GetUsers", ctx).Return(nil, expectedError).Once()

		// Act
		users, err := svc.GetUsers(ctx)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, users)
		mockRepo.AssertExpectations(t)
	})
}

func TestUserService_GetUserById(t *testing.T) {
	// Arrange
	mockRepo := new(mockUserRepo)
	svc := NewUserService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		userId := "1"
		expectedUser := &models.User{ID: "1", Username: "user1", Email: "user1@example.com"}
		mockRepo.On("GetUserByID", ctx, userId).Return(expectedUser, nil).Once()

		// Act
		user, err := svc.GetUserById(ctx, userId)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Not Found", func(t *testing.T) {
		userId := "999"
		mockRepo.On("GetUserByID", ctx, userId).Return(nil, nil).Once()

		// Act
		user, err := svc.GetUserById(ctx, userId)

		// Assert
		assert.NoError(t, err)
		assert.Nil(t, user)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		userId := "1"
		expectedError := errors.New("database error")
		mockRepo.On("GetUserByID", ctx, userId).Return(nil, expectedError).Once()

		// Act
		user, err := svc.GetUserById(ctx, userId)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, user)
		mockRepo.AssertExpectations(t)
	})
}

func TestUserService_CreateUser(t *testing.T) {
	// Arrange
	mockRepo := new(mockUserRepo)
	svc := NewUserService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		newUser := &models.User{
			Username: "newuser",
			Email:    "newuser@example.com",
			Password: "password123",
		}
		expectedUser := &models.User{
			ID:       "1",
			Username: "newuser",
			Email:    "newuser@example.com",
			Password: "password123",
		}
		mockRepo.On("CreateUser", ctx, newUser).Return(expectedUser, nil).Once()

		// Act
		user, err := svc.CreateUser(ctx, newUser)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		newUser := &models.User{
			Username: "newuser",
			Email:    "newuser@example.com",
		}
		expectedError := errors.New("database error")
		mockRepo.On("CreateUser", ctx, newUser).Return(nil, expectedError).Once()

		// Act
		user, err := svc.CreateUser(ctx, newUser)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, user)
		mockRepo.AssertExpectations(t)
	})
}

func TestUserService_UpdateUser(t *testing.T) {
	// Arrange
	mockRepo := new(mockUserRepo)
	svc := NewUserService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		updatedUser := models.User{
			ID:       "1",
			Username: "updateduser",
			Email:    "updated@example.com",
		}

		mockRepo.On("UpdateUser", ctx, updatedUser).Return(&updatedUser, nil).Once()

		// Act
		user, err := svc.UpdateUser(ctx, updatedUser)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, &updatedUser, user)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		updatedUser := models.User{
			ID:       "1",
			Username: "updateduser",
			Email:    "updated@example.com",
		}
		expectedError := errors.New("database error")
		mockRepo.On("UpdateUser", ctx, updatedUser).Return(nil, expectedError).Once()

		// Act
		user, err := svc.UpdateUser(ctx, updatedUser)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, user)
		mockRepo.AssertExpectations(t)
	})
}

func TestUserService_DeleteUser(t *testing.T) {
	// Arrange
	mockRepo := new(mockUserRepo)
	svc := NewUserService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		userId := "1"
		mockRepo.On("DeleteUser", ctx, userId).Return(nil).Once()

		// Act
		err := svc.DeleteUser(ctx, userId)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		userId := "1"
		expectedError := errors.New("database error")
		mockRepo.On("DeleteUser", ctx, userId).Return(expectedError).Once()

		// Act
		err := svc.DeleteUser(ctx, userId)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestUserService_GetUserByUsername(t *testing.T) {
	// Arrange
	mockRepo := new(mockUserRepo)
	svc := NewUserService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		username := "testuser"
		expectedUser := &models.User{ID: "1", Username: "testuser", Email: "test@example.com"}
		mockRepo.On("GetUserByUsername", ctx, username).Return(expectedUser, nil).Once()

		// Act
		user, err := svc.GetUserByUsername(ctx, username)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Not Found", func(t *testing.T) {
		username := "nonexistent"
		mockRepo.On("GetUserByUsername", ctx, username).Return(nil, nil).Once()

		// Act
		user, err := svc.GetUserByUsername(ctx, username)

		// Assert
		assert.NoError(t, err)
		assert.Nil(t, user)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		username := "testuser"
		expectedError := errors.New("database error")
		mockRepo.On("GetUserByUsername", ctx, username).Return(nil, expectedError).Once()

		// Act
		user, err := svc.GetUserByUsername(ctx, username)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, user)
		mockRepo.AssertExpectations(t)
	})
}
