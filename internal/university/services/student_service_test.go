package services

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"university_system/internal/domain/models"
)

type mockStudentRepo struct {
	mock.Mock
}

func (m *mockStudentRepo) GetStudents(ctx context.Context) ([]models.Student, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Student), args.Error(1)
}

func (m *mockStudentRepo) GetStudentById(ctx context.Context, id string) (*models.Student, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Student), args.Error(1)
}

func (m *mockStudentRepo) CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error) {
	args := m.Called(ctx, student)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Student), args.Error(1)
}

func (m *mockStudentRepo) UpdateStudent(ctx context.Context, student models.Student) (*models.Student, error) {
	args := m.Called(ctx, student)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Student), args.Error(1)
}

func (m *mockStudentRepo) DeleteStudent(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *mockStudentRepo) EnrollStudentToCourse(ctx context.Context, studentID, courseID string) error {
	args := m.Called(ctx, studentID, courseID)
	return args.Error(0)
}

func (m *mockStudentRepo) GetStudentCourses(ctx context.Context, studentID string) ([]models.Course, error) {
	args := m.Called(ctx, studentID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Course), args.Error(1)
}

func (m *mockStudentRepo) CreateUserWithRole(ctx context.Context, user models.User, role string) (string, error) {
	args := m.Called(ctx, user, role)
	return args.String(0), args.Error(1)
}

func TestStudentService_GetStudents(t *testing.T) {
	// Arrange
	mockRepo := new(mockStudentRepo)
	svc := NewStudentService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		expectedStudents := []models.Student{
			{StudentYear: 1, Faculty: "CS", CreatedAt: "2022-01-01", UpdatedAt: "2022-01-01", DeletedAt: nil, User: models.User{ID: "1"}},
			{StudentYear: 2, Faculty: "CS", CreatedAt: "2022-01-01", UpdatedAt: "2022-01-01", DeletedAt: nil, User: models.User{ID: "2"}},
		}
		mockRepo.On("GetStudents", ctx).Return(expectedStudents, nil).Once()

		// Act
		students, err := svc.GetStudents(ctx)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedStudents, students)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("database error")
		mockRepo.On("GetStudents", ctx).Return(nil, expectedError).Once()

		// Act
		students, err := svc.GetStudents(ctx)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, students)
		mockRepo.AssertExpectations(t)
	})
}

func TestStudentService_GetStudentById(t *testing.T) {
	// Arrange
	mockRepo := new(mockStudentRepo)
	svc := NewStudentService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		studentId := "1"
		expectedStudent := &models.Student{StudentYear: 1, Faculty: "CS", CreatedAt: "2022-01-01", UpdatedAt: "2022-01-01", DeletedAt: nil, User: models.User{ID: "1"}}
		mockRepo.On("GetStudentById", ctx, studentId).Return(expectedStudent, nil).Once()

		// Act
		student, err := svc.GetStudentById(ctx, studentId)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedStudent, student)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Not Found", func(t *testing.T) {
		studentId := "999"
		mockRepo.On("GetStudentById", ctx, studentId).Return(nil, nil).Once()

		// Act
		student, err := svc.GetStudentById(ctx, studentId)

		// Assert
		assert.NoError(t, err)
		assert.Nil(t, student)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		studentId := "1"
		expectedError := errors.New("database error")
		mockRepo.On("GetStudentById", ctx, studentId).Return(nil, expectedError).Once()

		// Act
		student, err := svc.GetStudentById(ctx, studentId)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, student)
		mockRepo.AssertExpectations(t)
	})
}

func TestStudentService_CreateStudent(t *testing.T) {
	// Arrange
	mockRepo := new(mockStudentRepo)
	svc := NewStudentService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		newStudent := &models.Student{
			StudentYear: 1, Faculty: "CS", CreatedAt: "2022-01-01", UpdatedAt: "2022-01-01", DeletedAt: nil, User: models.User{ID: "1"},
		}
		expectedStudent := &models.Student{
			StudentYear: 1, Faculty: "CS", CreatedAt: "2022-01-01", UpdatedAt: "2022-01-01", DeletedAt: nil, User: models.User{ID: "1"},
		}
		mockRepo.On("CreateStudent", ctx, newStudent).Return(expectedStudent, nil).Once()

		// Act
		student, err := svc.CreateStudent(ctx, newStudent)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedStudent, student)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		newStudent := &models.Student{
			StudentYear: 1, Faculty: "CS", CreatedAt: "2022-01-01", UpdatedAt: "2022-01-01", DeletedAt: nil, User: models.User{ID: "1"},
		}
		expectedError := errors.New("database error")
		mockRepo.On("CreateStudent", ctx, newStudent).Return(nil, expectedError).Once()

		// Act
		student, err := svc.CreateStudent(ctx, newStudent)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, student)
		mockRepo.AssertExpectations(t)
	})
}

func TestStudentService_UpdateStudent(t *testing.T) {
	// Arrange
	mockRepo := new(mockStudentRepo)
	svc := NewStudentService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		updatedStudent := models.Student{
			StudentYear: 1, Faculty: "CS", CreatedAt: "2022-01-01", UpdatedAt: "2022-01-01", DeletedAt: nil, User: models.User{ID: "1"},
		}
		mockRepo.On("UpdateStudent", ctx, updatedStudent).Return(&updatedStudent, nil).Once()

		// Act
		student, err := svc.UpdateStudent(ctx, updatedStudent)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, &updatedStudent, student)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		updatedStudent := models.Student{
			StudentYear: 1, Faculty: "CS", CreatedAt: "2022-01-01", UpdatedAt: "2022-01-01", DeletedAt: nil, User: models.User{ID: "1"},
		}
		expectedError := errors.New("database error")
		mockRepo.On("UpdateStudent", ctx, updatedStudent).Return(nil, expectedError).Once()

		// Act
		student, err := svc.UpdateStudent(ctx, updatedStudent)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, student)
		mockRepo.AssertExpectations(t)
	})
}

func TestStudentService_DeleteStudent(t *testing.T) {
	// Arrange
	mockRepo := new(mockStudentRepo)
	svc := NewStudentService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		studentId := "1"
		mockRepo.On("DeleteStudent", ctx, studentId).Return(nil).Once()

		// Act
		err := svc.DeleteStudent(ctx, studentId)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		studentId := "1"
		expectedError := errors.New("database error")
		mockRepo.On("DeleteStudent", ctx, studentId).Return(expectedError).Once()

		// Act
		err := svc.DeleteStudent(ctx, studentId)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestStudentService_EnrollStudentToCourse(t *testing.T) {
	// Arrange
	mockRepo := new(mockStudentRepo)
	svc := NewStudentService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		studentId := "1"
		courseId := "101"
		mockRepo.On("EnrollStudentToCourse", ctx, studentId, courseId).Return(nil).Once()

		// Act
		err := svc.EnrollStudentToCourse(ctx, studentId, courseId)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		studentId := "1"
		courseId := "101"
		expectedError := errors.New("database error")
		mockRepo.On("EnrollStudentToCourse", ctx, studentId, courseId).Return(expectedError).Once()

		// Act
		err := svc.EnrollStudentToCourse(ctx, studentId, courseId)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestStudentService_GetStudentCourses(t *testing.T) {
	// Arrange
	mockRepo := new(mockStudentRepo)
	svc := NewStudentService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		studentId := "1"
		expectedCourses := []models.Course{
			{ID: "101", Name: "Introduction to CS", Description: "Basic CS concepts"},
			{ID: "102", Name: "Advanced Algorithms", Description: "Complex algorithms study"},
		}
		mockRepo.On("GetStudentCourses", ctx, studentId).Return(expectedCourses, nil).Once()

		// Act
		courses, err := svc.GetStudentCourses(ctx, studentId)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedCourses, courses)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Empty Courses", func(t *testing.T) {
		studentId := "2"
		expectedCourses := []models.Course{}
		mockRepo.On("GetStudentCourses", ctx, studentId).Return(expectedCourses, nil).Once()

		// Act
		courses, err := svc.GetStudentCourses(ctx, studentId)

		// Assert
		assert.NoError(t, err)
		assert.Empty(t, courses)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		studentId := "1"
		expectedError := errors.New("database error")
		mockRepo.On("GetStudentCourses", ctx, studentId).Return(nil, expectedError).Once()

		// Act
		courses, err := svc.GetStudentCourses(ctx, studentId)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, courses)
		mockRepo.AssertExpectations(t)
	})
}
