package services

import (
	"context"
	"university_system/internal/domain/models"
	"university_system/internal/domain/repository"
)

type StudentService interface {
	GetStudents(ctx context.Context) ([]models.Student, error)
	GetStudentById(ctx context.Context, id string) (*models.Student, error)
	CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error)
	UpdateStudent(ctx context.Context, student models.Student) (*models.Student, error)
	DeleteStudent(ctx context.Context, id string) error
	EnrollStudentToCourse(ctx context.Context, studentID, courseID string) error
	GetStudentCourses(ctx context.Context, studentID string) ([]models.Course, error)
	CreateUserWithRole(ctx context.Context, user models.User, role string) (string, error)
}

type studentService struct {
	repo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) StudentService {
	return &studentService{repo: repo}
}

func (s *studentService) GetStudents(ctx context.Context) ([]models.Student, error) {
	return s.repo.GetStudents(ctx)
}

func (s *studentService) GetStudentById(ctx context.Context, id string) (*models.Student, error) {
	return s.repo.GetStudentById(ctx, id)
}

func (s *studentService) CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error) {
	return s.repo.CreateStudent(ctx, student)
}

func (s *studentService) UpdateStudent(ctx context.Context, student models.Student) (*models.Student, error) {
	return s.repo.UpdateStudent(ctx, student)
}

func (s *studentService) DeleteStudent(ctx context.Context, id string) error {
	return s.repo.DeleteStudent(ctx, id)
}

func (s *studentService) EnrollStudentToCourse(ctx context.Context, studentID, courseID string) error {
	return s.repo.EnrollStudentToCourse(ctx, studentID, courseID)
}

func (s *studentService) GetStudentCourses(ctx context.Context, studentID string) ([]models.Course, error) {
	return s.repo.GetStudentCourses(ctx, studentID)
}

// CreateUserWithRole создаёт пользователя с ролью и возвращает id
func (s *studentService) CreateUserWithRole(ctx context.Context, user models.User, role string) (string, error) {
	return s.repo.CreateUserWithRole(ctx, user, role)
}
