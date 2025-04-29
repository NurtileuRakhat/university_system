package services

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"university_system/internal/domain/models"
)

type mockGradeRepo struct {
	mock.Mock
}

func (m *mockGradeRepo) AddMark(ctx context.Context, mark *models.Mark, markType string) error {
	args := m.Called(ctx, mark, markType)
	return args.Error(0)
}

func (m *mockGradeRepo) GetStudentMarks(ctx context.Context, studentID string) ([]models.Mark, error) {
	args := m.Called(ctx, studentID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Mark), args.Error(1)
}

func (m *mockGradeRepo) GetCourseMarks(ctx context.Context, courseID string) ([]models.Mark, error) {
	args := m.Called(ctx, courseID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Mark), args.Error(1)
}

func (m *mockGradeRepo) IsTeacherOfCourse(ctx context.Context, teacherID string, courseID string) (bool, error) {
	args := m.Called(ctx, teacherID, courseID)
	return args.Bool(0), args.Error(1)
}

func TestGradeService_GetStudentMarks(t *testing.T) {
	// Arrange
	mockRepo := new(mockGradeRepo)
	svc := NewGradeService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		studentID := uint(101)
		expectedMarks := []models.Mark{
			{
				ID:                1,
				StudentID:         uint(101),
				CourseID:          uint(201),
				FirstAttestation:  25.5,
				SecondAttestation: 27.0,
				FinalMark:         38.0,
			},
			{
				ID:                2,
				StudentID:         uint(101),
				CourseID:          uint(202),
				FirstAttestation:  24.0,
				SecondAttestation: 26.0,
				FinalMark:         35.0,
			},
		}
		mockRepo.On("GetStudentMarks", ctx, fmt.Sprint(studentID)).Return(expectedMarks, nil).Once()

		// Act
		marks, err := svc.GetStudentMarks(ctx, fmt.Sprint(studentID))

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedMarks, marks)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Empty Marks", func(t *testing.T) {
		studentID := uint(102)
		expectedMarks := []models.Mark{}
		mockRepo.On("GetStudentMarks", ctx, fmt.Sprint(studentID)).Return(expectedMarks, nil).Once()

		// Act
		marks, err := svc.GetStudentMarks(ctx, fmt.Sprint(studentID))

		// Assert
		assert.NoError(t, err)
		assert.Empty(t, marks)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		studentID := uint(101)
		expectedError := errors.New("database error")
		mockRepo.On("GetStudentMarks", ctx, fmt.Sprint(studentID)).Return(nil, expectedError).Once()

		// Act
		marks, err := svc.GetStudentMarks(ctx, fmt.Sprint(studentID))

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, marks)
		mockRepo.AssertExpectations(t)
	})
}

func TestGradeService_GetCourseMarks(t *testing.T) {
	// Arrange
	mockRepo := new(mockGradeRepo)
	svc := NewGradeService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		courseID := uint(201)
		expectedMarks := []models.Mark{
			{
				ID:                1,
				StudentID:         uint(101),
				CourseID:          uint(201),
				FirstAttestation:  25.5,
				SecondAttestation: 27.0,
				FinalMark:         38.0,
			},
			{
				ID:                3,
				StudentID:         uint(102),
				CourseID:          uint(201),
				FirstAttestation:  28.5,
				SecondAttestation: 29.0,
				FinalMark:         39.0,
			},
		}
		mockRepo.On("GetCourseMarks", ctx, fmt.Sprint(courseID)).Return(expectedMarks, nil).Once()

		// Act
		marks, err := svc.GetCourseMarks(ctx, fmt.Sprint(courseID))

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedMarks, marks)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Empty Marks", func(t *testing.T) {
		courseID := uint(203)
		expectedMarks := []models.Mark{}
		mockRepo.On("GetCourseMarks", ctx, fmt.Sprint(courseID)).Return(expectedMarks, nil).Once()

		// Act
		marks, err := svc.GetCourseMarks(ctx, fmt.Sprint(courseID))

		// Assert
		assert.NoError(t, err)
		assert.Empty(t, marks)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		courseID := uint(201)
		expectedError := errors.New("database error")
		mockRepo.On("GetCourseMarks", ctx, fmt.Sprint(courseID)).Return(nil, expectedError).Once()

		// Act
		marks, err := svc.GetCourseMarks(ctx, fmt.Sprint(courseID))

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, marks)
		mockRepo.AssertExpectations(t)
	})
}
