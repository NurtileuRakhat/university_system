package repository

import (
	"context"
	"university_system/internal/domain/models"
)

type ManagerRepository interface {
	GetManagers(ctx context.Context) ([]models.Manager, error)
	GetManagerById(ctx context.Context, id string) (*models.Manager, error)
	CreateManager(ctx context.Context, manager *models.Manager) (*models.Manager, error)
	UpdateManager(ctx context.Context, manager models.Manager) (*models.Manager, error)
	DeleteManager(ctx context.Context, id string) error
	AssignTeacherToCourse(ctx context.Context, teacherID string, courseID string) error
	CreateUserWithRole(ctx context.Context, user models.User, role string) (string, error)
}
