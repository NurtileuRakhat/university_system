package repository

import (
	"errors"
	"gorm.io/gorm"
	"university_system/internal/university/models"
)

type ManagerRepository interface {
	GetManagers() ([]models.Manager, error)
	GetManagerById(ID string) (*models.Manager, error)
	CreateManager(manager models.Manager) (*models.Manager, error)
	UpdateManager(manager *models.Manager) (*models.Manager, error)
	DeleteManager(ID string) error
	GetManagerByUsername(username string) (*models.Manager, error)
	AssignTeacherToCourse(studentID, courseID string) error
}

type ManagerRepoImpl struct {
	DB      *gorm.DB
	manager models.Manager
}

func NewManagerRepository(db *gorm.DB) ManagerRepository {
	return &ManagerRepoImpl{DB: db}
}

func (r *ManagerRepoImpl) GetManagerById(ID string) (*models.Manager, error) {
	var manager models.Manager
	if err := r.DB.First(&manager, "id = ?", ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &manager, nil
}

func (r *ManagerRepoImpl) GetManagers() ([]models.Manager, error) {
	var managers []models.Manager
	if err := r.DB.Find(&managers).Error; err != nil {
		return nil, err
	}
	return managers, nil
}

func (r *ManagerRepoImpl) CreateManager(manager models.Manager) (*models.Manager, error) {
	user := models.User{
		Username:  manager.Username,
		Password:  manager.Password,
		Firstname: manager.Firstname,
		Lastname:  manager.Lastname,
		Email:     manager.Email,
		Role:      "Manager",
	}
	if err := r.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	manager.ID = user.ID
	if err := r.DB.Create(&manager).Error; err != nil {
		return nil, err
	}
	return &manager, nil
}

func (r *ManagerRepoImpl) UpdateManager(manager *models.Manager) (*models.Manager, error) {
	if err := r.DB.Save(&manager).Error; err != nil {
		return nil, err
	}
	var user models.User
	if err := r.DB.First(&user, "id = ?", manager.ID).Error; err != nil {
		return nil, err
	}
	user.Username = manager.Username
	user.Password = manager.Password
	user.Firstname = manager.Firstname
	user.Lastname = manager.Lastname
	user.Email = manager.Email

	if err := r.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return manager, nil
}

func (r *ManagerRepoImpl) DeleteManager(ID string) error {
	var manager models.Manager
	if err := r.DB.Delete(&models.User{}, "id = ?", ID).Error; err != nil {
		return err
	}
	if err := r.DB.First(&manager, ID).Error; err != nil {
		return err
	}
	r.DB.Delete(&manager)
	return nil
}

func (r *ManagerRepoImpl) GetManagerByUsername(username string) (*models.Manager, error) {
	var manager models.Manager
	if err := r.DB.First(&manager, "username = ?", username).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &manager, nil
}

func (r *ManagerRepoImpl) AssignTeacherToCourse(courseId string, teacherId string) error {
	var teacher models.Teacher
	var course models.Course

	if err := r.DB.First(&teacher, "id = ?", teacherId).Error; err != nil {
		return err
	}
	if err := r.DB.First(&course, "id = ?", courseId).Error; err != nil {
		return err
	}

	if err := r.DB.Model(&teacher).Association("Courses").Append(&course); err != nil {
		return err
	}

	return nil
}
