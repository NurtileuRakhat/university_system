package services

import (
	"context"
	"university_system/internal/domain/models"
	"university_system/internal/domain/repository"
)

type TeacherService interface {
	GetTeachers(ctx context.Context) ([]models.Teacher, error)
	GetTeacherById(ctx context.Context, id string) (*models.Teacher, error)
	UpdateTeacher(ctx context.Context, teacher models.Teacher) (*models.Teacher, error)
	DeleteTeacher(ctx context.Context, id string) error
	CreateTeacher(ctx context.Context, teacher *models.Teacher) (*models.Teacher, error)
	GetTeacherCourses(ctx context.Context, teacherID string) ([]models.Course, error)
	CreateUserWithRole(ctx context.Context, user models.User, role string) (string, error)
}

type teacherService struct {
	repo repository.TeacherRepository
}

func NewTeacherService(repo repository.TeacherRepository) TeacherService {
	return &teacherService{repo: repo}
}

// GetTeachers возвращает список всех преподавателей
func (s *teacherService) GetTeachers(ctx context.Context) ([]models.Teacher, error) {
	return s.repo.GetTeachers(ctx)
}

// GetTeacherById возвращает преподавателя по ID
func (s *teacherService) GetTeacherById(ctx context.Context, id string) (*models.Teacher, error) {
	return s.repo.GetTeacherById(ctx, id)
}

// UpdateTeacher обновляет информацию о преподавателе
func (s *teacherService) UpdateTeacher(ctx context.Context, teacher models.Teacher) (*models.Teacher, error) {
	return s.repo.UpdateTeacher(ctx, teacher)
}

// DeleteTeacher удаляет преподавателя по ID
func (s *teacherService) DeleteTeacher(ctx context.Context, id string) error {
	return s.repo.DeleteTeacher(ctx, id)
}

// CreateTeacher создает нового преподавателя
func (s *teacherService) CreateTeacher(ctx context.Context, teacher *models.Teacher) (*models.Teacher, error) {
	return s.repo.CreateTeacher(ctx, teacher)
}

// GetTeacherCourses возвращает список курсов, которые ведет преподаватель
func (s *teacherService) GetTeacherCourses(ctx context.Context, teacherID string) ([]models.Course, error) {
	return s.repo.GetTeacherCourses(ctx, teacherID)
}

// CreateUserWithRole создаёт пользователя с ролью и возвращает id
func (s *teacherService) CreateUserWithRole(ctx context.Context, user models.User, role string) (string, error) {
	return s.repo.CreateUserWithRole(ctx, user, role)
}