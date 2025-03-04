package repository

import (
	"errors"
	"gorm.io/gorm"
	"university_system/internal/university/models"
)

type TeacherRepository interface {
	GetTeachers() ([]models.Teacher, error)
	GetTeacherByID(int) (*models.Teacher, error)
	UpdateTeacher(teacher *models.Teacher) (*models.Teacher, error)
	DeleteTeacher(int) error
	CreateTeacher(teacher *models.Teacher) (*models.Teacher, error)
	GetTeacherByUsername(username string) (*models.Teacher, error)
	GetTeacherCourses(teacherID int) ([]models.Course, error)
}

type TeacherRepositoryImpl struct {
	DB *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) TeacherRepository {
	return &TeacherRepositoryImpl{DB: db}
}

func (t *TeacherRepositoryImpl) GetTeachers() ([]models.Teacher, error) {
	var teachers []models.Teacher
	if err := t.DB.Find(&teachers).Error; err != nil {
		return nil, err
	}
	return teachers, nil
}

func (t *TeacherRepositoryImpl) GetTeacherByID(id int) (*models.Teacher, error) {
	var teacher models.Teacher
	if err := t.DB.First(&teacher, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &teacher, nil
}

func (t *TeacherRepositoryImpl) UpdateTeacher(teacher *models.Teacher) (*models.Teacher, error) {
	if err := t.DB.Save(teacher).Error; err != nil {
		return nil, err
	}
	var user models.User
	if err := t.DB.First(&user, "id = ?", teacher.ID).Error; err != nil {
		return nil, err
	}
	user.Username = teacher.Username
	user.Password = teacher.Password
	user.Firstname = teacher.Firstname
	user.Lastname = teacher.Lastname
	user.Email = teacher.Email

	if err := t.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return teacher, nil
}

func (t *TeacherRepositoryImpl) DeleteTeacher(id int) error {
	if err := t.DB.Delete(&models.User{}, "id = ?", id).Error; err != nil {
		return err
	}
	if err := t.DB.Delete(&models.Teacher{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (t *TeacherRepositoryImpl) CreateTeacher(teacher *models.Teacher) (*models.Teacher, error) {
	user := models.User{
		Username:  teacher.Username,
		Password:  teacher.Password,
		Firstname: teacher.Firstname,
		Lastname:  teacher.Lastname,
		Email:     teacher.Email,
		Role:      "Teacher",
	}
	if err := t.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	teacher.ID = user.ID
	if err := t.DB.Create(teacher).Error; err != nil {
		return nil, err
	}
	return teacher, nil
}

func (t *TeacherRepositoryImpl) GetTeacherByUsername(username string) (*models.Teacher, error) {
	var teacher models.Teacher
	if err := t.DB.First(&teacher, "username = ?", username).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &teacher, nil
}

func (t *TeacherRepositoryImpl) GetTeacherCourses(teacherID int) ([]models.Course, error) {
	var teacher models.Teacher

	if err := t.DB.Preload("Courses").First(&teacher, "id = ?", teacherID).Error; err != nil {
		return nil, err
	}
	return teacher.Courses, nil
}
