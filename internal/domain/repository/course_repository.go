package repository

import (
	"context"
	"university_system/internal/domain/models"
)

type CourseRepository interface {
	CreateCourse(ctx context.Context, course *models.Course) (*models.Course, error)
	GetCourseByID(ctx context.Context, ID string) (*models.Course, error)
	GetAllCourses(ctx context.Context) ([]models.Course, error)
	UpdateCourse(ctx context.Context, course models.Course) (*models.Course, error)
	DeleteCourse(ctx context.Context, ID string) error
	GetCourseStudents(ctx context.Context, courseID string) ([]models.Student, error)
	GetCourseTeachers(ctx context.Context, courseID string) ([]models.Teacher, error)
}
