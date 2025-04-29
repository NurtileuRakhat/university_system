package repository

import (
	"context"
	"university_system/internal/domain/models"
)

type GradeRepository interface {
	GetStudentMarks(ctx context.Context, studentID string) ([]models.Mark, error)
	GetCourseMarks(ctx context.Context, courseID string) ([]models.Mark, error)
	AddMark(ctx context.Context, mark *models.Mark, markType string) error
	IsTeacherOfCourse(ctx context.Context, teacherID string, courseID string) (bool, error)
}
