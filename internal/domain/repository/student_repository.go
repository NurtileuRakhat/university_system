package repository

import (
	"context"
	"university_system/internal/domain/models"
)

type StudentRepository interface {
	GetStudents(ctx context.Context) ([]models.Student, error)
	GetStudentById(ctx context.Context, id string) (*models.Student, error)
	CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error)
	UpdateStudent(ctx context.Context, student models.Student) (*models.Student, error)
	DeleteStudent(ctx context.Context, id string) error
	EnrollStudentToCourse(ctx context.Context, studentID, courseID string) error
	GetStudentCourses(ctx context.Context, studentID string) ([]models.Course, error)
	CreateUserWithRole(ctx context.Context, user models.User, role string) (string, error)
}
