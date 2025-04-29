package services

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"university_system/internal/domain/models"
)

type mockTeacherRepo struct {
	mock.Mock
}

func (m *mockTeacherRepo) GetTeachers(ctx context.Context) ([]models.Teacher, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Teacher), args.Error(1)
}

func (m *mockTeacherRepo) GetTeacherById(ctx context.Context, id string) (*models.Teacher, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Teacher), args.Error(1)
}

func (m *mockTeacherRepo) CreateTeacher(ctx context.Context, teacher *models.Teacher) (*models.Teacher, error) {
	args := m.Called(ctx, teacher)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Teacher), args.Error(1)
}

func (m *mockTeacherRepo) UpdateTeacher(ctx context.Context, teacher models.Teacher) (*models.Teacher, error) {
	args := m.Called(ctx, teacher)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Teacher), args.Error(1)
}

func (m *mockTeacherRepo) DeleteTeacher(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *mockTeacherRepo) AssignTeacherToCourse(ctx context.Context, teacherID, courseID string) error {
	args := m.Called(ctx, teacherID, courseID)
	return args.Error(0)
}

func (m *mockTeacherRepo) GetTeacherCourses(ctx context.Context, teacherID string) ([]models.Course, error) {
	args := m.Called(ctx, teacherID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Course), args.Error(1)
}

func (m *mockTeacherRepo) CreateUserWithRole(ctx context.Context, user models.User, role string) (string, error) {
	args := m.Called(ctx, user, role)
	return args.String(0), args.Error(1)
}

func TestTeacherService_GetTeachers(t *testing.T) {
	// Arrange
	mockRepo := new(mockTeacherRepo)
	svc := NewTeacherService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		expectedTeachers := []models.Teacher{
			{User: models.User{ID: "1"}, Department: "Computer Science", Position: "Professor"},
			{User: models.User{ID: "2"}, Department: "Mathematics", Position: "Associate Professor"},
		}
		mockRepo.On("GetTeachers", ctx).Return(expectedTeachers, nil).Once()

		// Act
		teachers, err := svc.GetTeachers(ctx)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedTeachers, teachers)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("database error")
		mockRepo.On("GetTeachers", ctx).Return(nil, expectedError).Once()

		// Act
		teachers, err := svc.GetTeachers(ctx)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, teachers)
		mockRepo.AssertExpectations(t)
	})
}

func TestTeacherService_GetTeacherById(t *testing.T) {
	// Arrange
	mockRepo := new(mockTeacherRepo)
	svc := NewTeacherService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		teacherID := "1"
		expectedTeacher := &models.Teacher{User: models.User{ID: "1"}, Department: "Computer Science", Position: "Professor"}
		mockRepo.On("GetTeacherById", ctx, teacherID).Return(expectedTeacher, nil).Once()

		// Act
		teacher, err := svc.GetTeacherById(ctx, teacherID)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedTeacher, teacher)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Not Found", func(t *testing.T) {
		teacherID := "999"
		mockRepo.On("GetTeacherById", ctx, teacherID).Return(nil, nil).Once()

		// Act
		teacher, err := svc.GetTeacherById(ctx, teacherID)

		// Assert
		assert.NoError(t, err)
		assert.Nil(t, teacher)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		teacherID := "1"
		expectedError := errors.New("database error")
		mockRepo.On("GetTeacherById", ctx, teacherID).Return(nil, expectedError).Once()

		// Act
		teacher, err := svc.GetTeacherById(ctx, teacherID)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, teacher)
		mockRepo.AssertExpectations(t)
	})
}

func TestTeacherService_CreateTeacher(t *testing.T) {
	// Arrange
	mockRepo := new(mockTeacherRepo)
	svc := NewTeacherService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		newTeacher := &models.Teacher{
			User:       models.User{ID: "1"},
			Department: "Computer Science",
			Position:   "Professor",
		}
		expectedTeacher := &models.Teacher{
			User:       models.User{ID: "1"},
			Department: "Computer Science",
			Position:   "Professor",
		}
		mockRepo.On("CreateTeacher", ctx, newTeacher).Return(expectedTeacher, nil).Once()

		// Act
		teacher, err := svc.CreateTeacher(ctx, newTeacher)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedTeacher, teacher)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		newTeacher := &models.Teacher{
			User:       models.User{ID: "1"},
			Department: "Computer Science",
			Position:   "Professor",
		}
		expectedError := errors.New("database error")
		mockRepo.On("CreateTeacher", ctx, newTeacher).Return(nil, expectedError).Once()

		// Act
		teacher, err := svc.CreateTeacher(ctx, newTeacher)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, teacher)
		mockRepo.AssertExpectations(t)
	})
}

func TestTeacherService_UpdateTeacher(t *testing.T) {
	// Arrange
	mockRepo := new(mockTeacherRepo)
	svc := NewTeacherService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		updatedTeacher := models.Teacher{
			User:       models.User{ID: "1"},
			Department: "Computer Science",
			Position:   "Senior Professor", // Изменена должность
		}
		mockRepo.On("UpdateTeacher", ctx, updatedTeacher).Return(&updatedTeacher, nil).Once()

		// Act
		teacher, err := svc.UpdateTeacher(ctx, updatedTeacher)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, &updatedTeacher, teacher)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		updatedTeacher := models.Teacher{
			User:       models.User{ID: "1"},
			Department: "Computer Science",
			Position:   "Senior Professor",
		}
		expectedError := errors.New("database error")
		mockRepo.On("UpdateTeacher", ctx, updatedTeacher).Return(nil, expectedError).Once()

		// Act
		teacher, err := svc.UpdateTeacher(ctx, updatedTeacher)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, teacher)
		mockRepo.AssertExpectations(t)
	})
}

func TestTeacherService_DeleteTeacher(t *testing.T) {
	// Arrange
	mockRepo := new(mockTeacherRepo)
	svc := NewTeacherService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		teacherID := "1"
		mockRepo.On("DeleteTeacher", ctx, teacherID).Return(nil).Once()

		// Act
		err := svc.DeleteTeacher(ctx, teacherID)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		teacherID := "1"
		expectedError := errors.New("database error")
		mockRepo.On("DeleteTeacher", ctx, teacherID).Return(expectedError).Once()

		// Act
		err := svc.DeleteTeacher(ctx, teacherID)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestTeacherService_GetTeacherCourses(t *testing.T) {
	// Arrange
	mockRepo := new(mockTeacherRepo)
	svc := NewTeacherService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		teacherID := "1"
		expectedCourses := []models.Course{
			{ID: "101", Name: "Advanced Computer Science", Description: "Advanced CS concepts"},
			{ID: "102", Name: "Operating Systems", Description: "Study of operating systems"},
		}
		mockRepo.On("GetTeacherCourses", ctx, teacherID).Return(expectedCourses, nil).Once()

		// Act
		courses, err := svc.GetTeacherCourses(ctx, teacherID)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedCourses, courses)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Empty Courses", func(t *testing.T) {
		teacherID := "2"
		expectedCourses := []models.Course{}
		mockRepo.On("GetTeacherCourses", ctx, teacherID).Return(expectedCourses, nil).Once()

		// Act
		courses, err := svc.GetTeacherCourses(ctx, teacherID)

		// Assert
		assert.NoError(t, err)
		assert.Empty(t, courses)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		teacherID := "1"
		expectedError := errors.New("database error")
		mockRepo.On("GetTeacherCourses", ctx, teacherID).Return(nil, expectedError).Once()

		// Act
		courses, err := svc.GetTeacherCourses(ctx, teacherID)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, courses)
		mockRepo.AssertExpectations(t)
	})
}
