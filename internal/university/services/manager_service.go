package services

import (
	"context"
	"university_system/internal/domain/models"
	"university_system/internal/domain/repository"
)

type ManagerService interface {
	GetManagers(ctx context.Context) ([]models.Manager, error)
	GetManagerById(ctx context.Context, ID string) (*models.Manager, error)
	UpdateManager(ctx context.Context, manager models.Manager) (*models.Manager, error)
	DeleteManager(ctx context.Context, ID string) error
	CreateManager(ctx context.Context, manager *models.Manager) (*models.Manager, error)
	CreateUserWithRole(ctx context.Context, user models.User, role string) (string, error)
	AssignTeacherToCourse(ctx context.Context, teacherID, courseID string) error
}

type managerService struct {
	repo repository.ManagerRepository
}

func NewManagerService(repo repository.ManagerRepository) ManagerService {
	return &managerService{repo: repo}
}

func (s *managerService) GetManagers(ctx context.Context) ([]models.Manager, error) {
	return s.repo.GetManagers(ctx)
}

func (s *managerService) GetManagerById(ctx context.Context, ID string) (*models.Manager, error) {
	return s.repo.GetManagerById(ctx, ID)
}

func (s *managerService) UpdateManager(ctx context.Context, manager models.Manager) (*models.Manager, error) {
	return s.repo.UpdateManager(ctx, manager)
}

func (s *managerService) DeleteManager(ctx context.Context, ID string) error {
	return s.repo.DeleteManager(ctx, ID)
}

func (s *managerService) CreateManager(ctx context.Context, manager *models.Manager) (*models.Manager, error) {
	return s.repo.CreateManager(ctx, manager)
}

func (s *managerService) CreateUserWithRole(ctx context.Context, user models.User, role string) (string, error) {
	return s.repo.CreateUserWithRole(ctx, user, role)
}

func (s *managerService) AssignTeacherToCourse(ctx context.Context, teacherID, courseID string) error {
	return s.repo.AssignTeacherToCourse(ctx, teacherID, courseID)
}
