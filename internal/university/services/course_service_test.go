package services

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"university_system/internal/domain/models"
)

type mockCourseRepo struct {
	mock.Mock
}

func (m *mockCourseRepo) CreateCourse(ctx context.Context, course *models.Course) (*models.Course, error) {
	args := m.Called(ctx, course)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Course), args.Error(1)
}

func (m *mockCourseRepo) GetCourseByID(ctx context.Context, ID string) (*models.Course, error) {
	args := m.Called(ctx, ID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Course), args.Error(1)
}

func (m *mockCourseRepo) GetAllCourses(ctx context.Context) ([]models.Course, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Course), args.Error(1)
}

func (m *mockCourseRepo) UpdateCourse(ctx context.Context, course models.Course) (*models.Course, error) {
	args := m.Called(ctx, course)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Course), args.Error(1)
}

func (m *mockCourseRepo) DeleteCourse(ctx context.Context, ID string) error {
	args := m.Called(ctx, ID)
	return args.Error(0)
}

func (m *mockCourseRepo) GetCourseStudents(ctx context.Context, courseID string) ([]models.Student, error) {
	args := m.Called(ctx, courseID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Student), args.Error(1)
}

func (m *mockCourseRepo) GetCourseTeachers(ctx context.Context, courseID string) ([]models.Teacher, error) {
	args := m.Called(ctx, courseID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Teacher), args.Error(1)
}

func TestCourseService_CreateCourse(t *testing.T) {
	// Arrange
	mockRepo := new(mockCourseRepo)
	svc := NewCourseService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		newCourse := &models.Course{
			Name:        "Data Structures",
			Description: "Course about data structures",
		}
		expectedCourse := &models.Course{
			ID:          "1",
			Name:        "Data Structures",
			Description: "Course about data structures",
		}
		mockRepo.On("CreateCourse", ctx, newCourse).Return(expectedCourse, nil).Once()

		// Act
		course, err := svc.CreateCourse(ctx, newCourse)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedCourse, course)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		newCourse := &models.Course{
			Name:        "Data Structures",
			Description: "Course about data structures",
		}
		expectedError := errors.New("database error")
		mockRepo.On("CreateCourse", ctx, newCourse).Return(nil, expectedError).Once()

		// Act
		course, err := svc.CreateCourse(ctx, newCourse)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, course)
		mockRepo.AssertExpectations(t)
	})
}

func TestCourseService_GetAllCourses(t *testing.T) {
	// Arrange
	mockRepo := new(mockCourseRepo)
	svc := NewCourseService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		expectedCourses := []models.Course{
			{ID: "1", Name: "Data Structures", Description: "Course about data structures"},
			{ID: "2", Name: "Algorithms", Description: "Course about algorithms"},
		}
		mockRepo.On("GetAllCourses", ctx).Return(expectedCourses, nil).Once()

		// Act
		courses, err := svc.GetAllCourses(ctx)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedCourses, courses)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Empty", func(t *testing.T) {
		expectedCourses := []models.Course{}
		mockRepo.On("GetAllCourses", ctx).Return(expectedCourses, nil).Once()

		// Act
		courses, err := svc.GetAllCourses(ctx)

		// Assert
		assert.NoError(t, err)
		assert.Empty(t, courses)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("database error")
		mockRepo.On("GetAllCourses", ctx).Return(nil, expectedError).Once()

		// Act
		courses, err := svc.GetAllCourses(ctx)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, courses)
		mockRepo.AssertExpectations(t)
	})
}

func TestCourseService_GetCourseByID(t *testing.T) {
	// Arrange
	mockRepo := new(mockCourseRepo)
	svc := NewCourseService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		courseID := "1"
		expectedCourse := &models.Course{
			ID:          "1",
			Name:        "Data Structures",
			Description: "Course about data structures",
		}
		mockRepo.On("GetCourseByID", ctx, courseID).Return(expectedCourse, nil).Once()

		// Act
		course, err := svc.GetCourseByID(ctx, courseID)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedCourse, course)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Not Found", func(t *testing.T) {
		courseID := "999"
		mockRepo.On("GetCourseByID", ctx, courseID).Return(nil, nil).Once()

		// Act
		course, err := svc.GetCourseByID(ctx, courseID)

		// Assert
		assert.NoError(t, err)
		assert.Nil(t, course)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		courseID := "1"
		expectedError := errors.New("database error")
		mockRepo.On("GetCourseByID", ctx, courseID).Return(nil, expectedError).Once()

		// Act
		course, err := svc.GetCourseByID(ctx, courseID)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, course)
		mockRepo.AssertExpectations(t)
	})
}

func TestCourseService_UpdateCourse(t *testing.T) {
	// Arrange
	mockRepo := new(mockCourseRepo)
	svc := NewCourseService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		updatedCourse := models.Course{
			ID:          "1",
			Name:        "Updated Data Structures",
			Description: "Updated course description",
		}
		mockRepo.On("UpdateCourse", ctx, updatedCourse).Return(&updatedCourse, nil).Once()

		// Act
		course, err := svc.UpdateCourse(ctx, updatedCourse)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, &updatedCourse, course)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		updatedCourse := models.Course{
			ID:          "1",
			Name:        "Updated Data Structures",
			Description: "Updated course description",
		}
		expectedError := errors.New("database error")
		mockRepo.On("UpdateCourse", ctx, updatedCourse).Return(nil, expectedError).Once()

		// Act
		course, err := svc.UpdateCourse(ctx, updatedCourse)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, course)
		mockRepo.AssertExpectations(t)
	})
}

func TestCourseService_DeleteCourse(t *testing.T) {
	// Arrange
	mockRepo := new(mockCourseRepo)
	svc := NewCourseService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		courseID := "1"
		mockRepo.On("DeleteCourse", ctx, courseID).Return(nil).Once()

		// Act
		err := svc.DeleteCourse(ctx, courseID)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		courseID := "1"
		expectedError := errors.New("database error")
		mockRepo.On("DeleteCourse", ctx, courseID).Return(expectedError).Once()

		// Act
		err := svc.DeleteCourse(ctx, courseID)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestCourseService_GetCourseStudents(t *testing.T) {
	// Arrange
	mockRepo := new(mockCourseRepo)
	svc := NewCourseService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		courseID := "1"
		expectedStudents := []models.Student{
			{StudentYear: 1, Faculty: "CS"},
			{StudentYear: 2, Faculty: "CS"},
		}
		mockRepo.On("GetCourseStudents", ctx, courseID).Return(expectedStudents, nil).Once()

		// Act
		students, err := svc.GetCourseStudents(ctx, courseID)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedStudents, students)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Empty Students", func(t *testing.T) {
		courseID := "2"
		expectedStudents := []models.Student{}
		mockRepo.On("GetCourseStudents", ctx, courseID).Return(expectedStudents, nil).Once()

		// Act
		students, err := svc.GetCourseStudents(ctx, courseID)

		// Assert
		assert.NoError(t, err)
		assert.Empty(t, students)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		courseID := "1"
		expectedError := errors.New("database error")
		mockRepo.On("GetCourseStudents", ctx, courseID).Return(nil, expectedError).Once()

		// Act
		students, err := svc.GetCourseStudents(ctx, courseID)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, students)
		mockRepo.AssertExpectations(t)
	})
}

func TestCourseService_GetCourseTeachers(t *testing.T) {
	// Arrange
	mockRepo := new(mockCourseRepo)
	svc := NewCourseService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		courseID := "1"
		expectedTeachers := []models.Teacher{
			{User: models.User{ID: "201"}, Department: "CS", Position: "Professor"},
			{User: models.User{ID: "202"}, Department: "CS", Position: "Assistant"},
		}
		mockRepo.On("GetCourseTeachers", ctx, courseID).Return(expectedTeachers, nil).Once()

		// Act
		teachers, err := svc.GetCourseTeachers(ctx, courseID)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedTeachers, teachers)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Empty Teachers", func(t *testing.T) {
		courseID := "2"
		expectedTeachers := []models.Teacher{}
		mockRepo.On("GetCourseTeachers", ctx, courseID).Return(expectedTeachers, nil).Once()

		// Act
		teachers, err := svc.GetCourseTeachers(ctx, courseID)

		// Assert
		assert.NoError(t, err)
		assert.Empty(t, teachers)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		courseID := "1"
		expectedError := errors.New("database error")
		mockRepo.On("GetCourseTeachers", ctx, courseID).Return(nil, expectedError).Once()

		// Act
		teachers, err := svc.GetCourseTeachers(ctx, courseID)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, teachers)
		mockRepo.AssertExpectations(t)
	})
}
