package services

import (
	"university_system/internal/university/repository"
)

type StudentService struct {
	Repo *repository.StudentRepository
}

func NewStudentService(repo *repository.StudentRepository) *StudentService {
	return &StudentService{Repo: repo}
}

//func (s *StudentService) RegisterStudent(student *models.Student) error {
//	return s.Repo.CreateStudent(student)
//}
//
//func (s *StudentService) GetStudent(id uint) (*models.Student, error) {
//	return s.Repo.GetStudentByID(id)
//}
//
//func (s *StudentService) EnrollStudentToCourse(studentID, courseID uint) error {
//	return s.Repo.EnrollToCourse(studentID, courseID)
//}
