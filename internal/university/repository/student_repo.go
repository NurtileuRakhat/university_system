package repository

import (
	"gorm.io/gorm"
	"university_system/internal/university/models"
)

type StudentRepository interface {
	GetStudents() ([]models.Student, error)
	GetStudentsById(ID string) (*models.Student, error)
	UpdateStudent(student *models.Student) (*models.Student, error)
	DeleteStudent(ID string) error
	CreateStudent(student *models.Student) (*models.Student, error)
	GetStudentByUsername(username string) (*models.Student, error)
	GetStudentCourses(studentID string) ([]models.Course, error)
	EnrollStudentToCourse(studentID, courseID string) error
}

type StudentRepositoryImpl struct {
	DB *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &StudentRepositoryImpl{DB: db}
}

func (s *StudentRepositoryImpl) GetStudents() ([]models.Student, error) {
	var students []models.Student
	if err := s.DB.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (s *StudentRepositoryImpl) GetStudentsById(ID string) (*models.Student, error) {
	var student models.Student
	if err := s.DB.Where("id = ?", ID).First(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (s *StudentRepositoryImpl) UpdateStudent(student *models.Student) (*models.Student, error) {
	if err := s.DB.Save(student).Error; err != nil {
		return nil, err
	}
	var user models.User
	if err := s.DB.First(&user, "id = ?", student.ID).Error; err != nil {
		return nil, err
	}
	user.Username = student.Username
	user.Password = student.Password
	user.Firstname = student.Firstname
	user.Lastname = student.Lastname
	user.Email = student.Email

	if err := s.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return student, nil
}

func (s *StudentRepositoryImpl) DeleteStudent(ID string) error {
	if err := s.DB.Delete(&models.User{}, "id = ?", ID).Error; err != nil {
		return err
	}
	if err := s.DB.Delete(&models.Student{}, "id = ?", ID).Error; err != nil {
		return err
	}
	return nil
}

func (s *StudentRepositoryImpl) CreateStudent(student *models.Student) (*models.Student, error) {
	user := models.User{
		Username:  student.Username,
		Password:  student.Password,
		Firstname: student.Firstname,
		Lastname:  student.Lastname,
		Email:     student.Email,
		Role:      "Student",
	}
	if err := s.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	student.ID = user.ID
	if err := s.DB.Create(student).Error; err != nil {
		return nil, err
	}
	return student, nil
}

func (s *StudentRepositoryImpl) GetStudentByUsername(username string) (*models.Student, error) {
	var student models.Student
	if err := s.DB.Where("username = ?", username).First(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (s *StudentRepositoryImpl) EnrollStudentToCourse(studentID, courseID string) error {
	var student models.Student
	var course models.Course

	if err := s.DB.First(&student, "id = ?", studentID).Error; err != nil {
		return err
	}
	if err := s.DB.First(&course, "id = ?", courseID).Error; err != nil {
		return err
	}

	if err := s.DB.Model(&student).Association("Courses").Append(&course); err != nil {
		return err
	}

	return nil
}

func (s *StudentRepositoryImpl) GetStudentCourses(studentID string) ([]models.Course, error) {
	var student models.Student
	if err := s.DB.Preload("Courses").First(&student, "id = ?", studentID).Error; err != nil {
		return nil, err
	}
	return student.Courses, nil
}
