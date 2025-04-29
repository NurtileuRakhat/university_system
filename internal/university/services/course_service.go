package services

import (
	"context"
	"university_system/internal/domain/models"
	"university_system/internal/domain/repository"
)

type CourseService interface {
	CreateCourse(ctx context.Context, course *models.Course) (*models.Course, error)
	GetAllCourses(ctx context.Context) ([]models.Course, error)
	GetCourseByID(ctx context.Context, id string) (*models.Course, error)
	UpdateCourse(ctx context.Context, course models.Course) (*models.Course, error)
	DeleteCourse(ctx context.Context, id string) error
	GetCourseStudents(ctx context.Context, courseID string) ([]models.Student, error)
	GetCourseTeachers(ctx context.Context, courseID string) ([]models.Teacher, error)
}

type courseService struct {
	repo repository.CourseRepository
}

func NewCourseService(repo repository.CourseRepository) CourseService {
	return &courseService{repo: repo}
}

func (s *courseService) CreateCourse(ctx context.Context, course *models.Course) (*models.Course, error) {
	return s.repo.CreateCourse(ctx, course)
}

func (s *courseService) GetAllCourses(ctx context.Context) ([]models.Course, error) {
	return s.repo.GetAllCourses(ctx)
}

func (s *courseService) GetCourseByID(ctx context.Context, id string) (*models.Course, error) {
	return s.repo.GetCourseByID(ctx, id)
}

func (s *courseService) UpdateCourse(ctx context.Context, course models.Course) (*models.Course, error) {
	return s.repo.UpdateCourse(ctx, course)
}

func (s *courseService) DeleteCourse(ctx context.Context, id string) error {
	return s.repo.DeleteCourse(ctx, id)
}

func (s *courseService) GetCourseStudents(ctx context.Context, courseID string) ([]models.Student, error) {
	return s.repo.GetCourseStudents(ctx, courseID)
}

func (s *courseService) GetCourseTeachers(ctx context.Context, courseID string) ([]models.Teacher, error) {
	return s.repo.GetCourseTeachers(ctx, courseID)
}