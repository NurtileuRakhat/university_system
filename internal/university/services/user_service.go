package services

import (
	"context"
	"university_system/internal/domain/models"
	"university_system/internal/domain/repository"
)

type UserService interface {
	GetUsers(ctx context.Context) ([]models.User, error)
	GetUserById(ctx context.Context, ID string) (*models.User, error)
	UpdateUser(ctx context.Context, user models.User) (*models.User, error)
	DeleteUser(ctx context.Context, ID string) error
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUsers(ctx context.Context) ([]models.User, error) {
	return s.repo.GetUsers(ctx)
}

func (s *userService) GetUserById(ctx context.Context, ID string) (*models.User, error) {
	return s.repo.GetUserByID(ctx, ID)
}

func (s *userService) UpdateUser(ctx context.Context, user models.User) (*models.User, error) {
	return s.repo.UpdateUser(ctx, user)
}

func (s *userService) DeleteUser(ctx context.Context, ID string) error {
	return s.repo.DeleteUser(ctx, ID)
}

func (s *userService) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	return s.repo.CreateUser(ctx, user)
}

func (s *userService) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	return s.repo.GetUserByUsername(ctx, username)
}
