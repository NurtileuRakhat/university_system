package repository

import (
	"errors"
	"gorm.io/gorm"
	"university_system/internal/university/models"
)

type UserRepository interface {
	GetUsers() ([]models.User, error)
	GetUserById(ID string) (*models.User, error)
	UpdateUser(user models.User) (*models.User, error)
	DeleteUser(ID string) error
	CreateUser(user *models.User) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
}

type UserRepositoryImpl struct {
	DB   *gorm.DB
	user models.User
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) GetUsers() ([]models.User, error) {
	var users []models.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepositoryImpl) GetUserById(ID string) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, "id = ?", ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Если пользователь не найден
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) UpdateUser(user models.User) (*models.User, error) {
	if err := r.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) DeleteUser(ID string) error {
	if err := r.DB.Delete(&models.User{}, "id = ?", ID).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) CreateUser(user *models.User) (*models.User, error) {
	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, "username = ?", username).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
