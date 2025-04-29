package services

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"university_system/internal/domain/models"
)

type mockManagerRepo struct {
	mock.Mock
}

func (m *mockManagerRepo) GetManagers(ctx context.Context) ([]models.Manager, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Manager), args.Error(1)
}

func (m *mockManagerRepo) GetManagerById(ctx context.Context, id string) (*models.Manager, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Manager), args.Error(1)
}

func (m *mockManagerRepo) CreateManager(ctx context.Context, manager *models.Manager) (*models.Manager, error) {
	args := m.Called(ctx, manager)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Manager), args.Error(1)
}

func (m *mockManagerRepo) UpdateManager(ctx context.Context, manager models.Manager) (*models.Manager, error) {
	args := m.Called(ctx, manager)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Manager), args.Error(1)
}

func (m *mockManagerRepo) DeleteManager(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *mockManagerRepo) AssignTeacherToCourse(ctx context.Context, teacherID, courseID string) error {
	args := m.Called(ctx, teacherID, courseID)
	return args.Error(0)
}

// Исправить CreateUserWithRole: принимать models.User, возвращать (string, error)
func (m *mockManagerRepo) CreateUserWithRole(ctx context.Context, user models.User, role string) (string, error) {
	args := m.Called(ctx, user, role)
	return args.String(0), args.Error(1)
}

func TestManagerService_GetManagers(t *testing.T) {
	// Arrange
	mockRepo := new(mockManagerRepo)
	svc := NewManagerService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		expectedManagers := []models.Manager{
			{User: models.User{ID: "1"}},
			{User: models.User{ID: "2"}},
		}
		mockRepo.On("GetManagers", ctx).Return(expectedManagers, nil).Once()

		// Act
		managers, err := svc.GetManagers(ctx)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedManagers, managers)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("database error")
		mockRepo.On("GetManagers", ctx).Return(nil, expectedError).Once()

		// Act
		managers, err := svc.GetManagers(ctx)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, managers)
		mockRepo.AssertExpectations(t)
	})
}

func TestManagerService_GetManagerById(t *testing.T) {
	// Arrange
	mockRepo := new(mockManagerRepo)
	svc := NewManagerService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		managerID := "1"
		expectedManager := &models.Manager{User: models.User{ID: "1"}}
		mockRepo.On("GetManagerById", ctx, managerID).Return(expectedManager, nil).Once()

		// Act
		manager, err := svc.GetManagerById(ctx, managerID)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedManager, manager)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Not Found", func(t *testing.T) {
		managerID := "999"
		mockRepo.On("GetManagerById", ctx, managerID).Return(nil, nil).Once()

		// Act
		manager, err := svc.GetManagerById(ctx, managerID)

		// Assert
		assert.NoError(t, err)
		assert.Nil(t, manager)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		managerID := "1"
		expectedError := errors.New("database error")
		mockRepo.On("GetManagerById", ctx, managerID).Return(nil, expectedError).Once()

		// Act
		manager, err := svc.GetManagerById(ctx, managerID)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, manager)
		mockRepo.AssertExpectations(t)
	})
}

func TestManagerService_CreateManager(t *testing.T) {
	// Arrange
	mockRepo := new(mockManagerRepo)
	svc := NewManagerService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		newManager := &models.Manager{
			User: models.User{ID: "1"},
		}
		expectedManager := &models.Manager{
			User: models.User{ID: "1"},
		}
		mockRepo.On("CreateManager", ctx, newManager).Return(expectedManager, nil).Once()

		// Act
		manager, err := svc.CreateManager(ctx, newManager)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedManager, manager)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		newManager := &models.Manager{
			User: models.User{ID: "1"},
		}
		expectedError := errors.New("database error")
		mockRepo.On("CreateManager", ctx, newManager).Return(nil, expectedError).Once()

		// Act
		manager, err := svc.CreateManager(ctx, newManager)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, manager)
		mockRepo.AssertExpectations(t)
	})
}

func TestManagerService_UpdateManager(t *testing.T) {
	// Arrange
	mockRepo := new(mockManagerRepo)
	svc := NewManagerService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		updatedManager := models.Manager{
			User: models.User{ID: "1"},
		}
		mockRepo.On("UpdateManager", ctx, updatedManager).Return(&updatedManager, nil).Once()

		// Act
		manager, err := svc.UpdateManager(ctx, updatedManager)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, &updatedManager, manager)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		updatedManager := models.Manager{
			User: models.User{ID: "1"},
		}
		expectedError := errors.New("database error")
		mockRepo.On("UpdateManager", ctx, updatedManager).Return(nil, expectedError).Once()

		// Act
		manager, err := svc.UpdateManager(ctx, updatedManager)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, manager)
		mockRepo.AssertExpectations(t)
	})
}

func TestManagerService_DeleteManager(t *testing.T) {
	// Arrange
	mockRepo := new(mockManagerRepo)
	svc := NewManagerService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		managerID := "1"
		mockRepo.On("DeleteManager", ctx, managerID).Return(nil).Once()

		// Act
		err := svc.DeleteManager(ctx, managerID)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		managerID := "1"
		expectedError := errors.New("database error")
		mockRepo.On("DeleteManager", ctx, managerID).Return(expectedError).Once()

		// Act
		err := svc.DeleteManager(ctx, managerID)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		mockRepo.AssertExpectations(t)
	})
}
