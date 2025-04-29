package repository

import (
	"context"
	"university_system/internal/domain/models"
)

type TeacherRepository interface {
	GetTeachers(ctx context.Context) ([]models.Teacher, error)
	GetTeacherById(ctx context.Context, id string) (*models.Teacher, error)
	CreateTeacher(ctx context.Context, teacher *models.Teacher) (*models.Teacher, error)
	UpdateTeacher(ctx context.Context, teacher models.Teacher) (*models.Teacher, error)
	DeleteTeacher(ctx context.Context, id string) error
	AssignTeacherToCourse(ctx context.Context, teacherID, courseID string) error
	GetTeacherCourses(ctx context.Context, teacherID string) ([]models.Course, error)
	CreateUserWithRole(ctx context.Context, user models.User, role string) (string, error)
}
