package services

import (
	"context"
	"university_system/internal/domain/models"
	"university_system/internal/domain/repository"
)

type GradeService interface {
	GetStudentMarks(ctx context.Context, studentID string) ([]models.Mark, error)
	GetCourseMarks(ctx context.Context, courseID string) ([]models.Mark, error)
}

type gradeService struct {
	repo repository.GradeRepository
}

func NewGradeService(repo repository.GradeRepository) GradeService {
	return &gradeService{repo: repo}
}

func (s *gradeService) GetStudentMarks(ctx context.Context, studentID string) ([]models.Mark, error) {
	return s.repo.GetStudentMarks(ctx, studentID)
}

func (s *gradeService) GetCourseMarks(ctx context.Context, courseID string) ([]models.Mark, error) {
	return s.repo.GetCourseMarks(ctx, courseID)
}
