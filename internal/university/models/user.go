package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type LoginRequest struct {
	Username string `json:"username" example:"user123"`
	Password string `json:"password" example:"mypassword"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// ErrorResponse
// @Description Структура ошибки API
type ErrorResponse struct {
	// Сообщение об ошибке
	// @example "Invalid request"
	Error   string `json:"error"`
	Message string `json:"message"`
}

// User представляет пользователя системы
// @Description Пользователь системы с уникальными данными
// @name User
// @swagger:model User
type User struct {
	// Встроенные поля GORM
	gorm.Model `gorm:"embedded"`

	// Уникальное имя пользователя
	Username string `gorm:"column:username;unique;not null" json:"username" binding:"required,min=3"`

	// Имя пользователя
	Firstname string `gorm:"column:firstname;not null" json:"firstname" binding:"required"`

	// Фамилия пользователя
	Lastname string `gorm:"column:lastname;not null" json:"lastname" binding:"required"`

	// Email пользователя (уникальный)
	Email string `gorm:"column:email;unique;not null" json:"email" binding:"required,email"`

	// Пароль пользователя
	Password string `gorm:"column:password;not null" json:"password" binding:"required,min=4"`

	// Дата рождения
	Birthdate time.Time `gorm:"column:birthdate" json:"birthdate" binding:"omitempty"`

	// Роль пользователя (админ, студент, учитель и т. д.)
	Role string `gorm:"column:role" json:"role" binding:"omitempty"`
}

func (user *User) CheckPassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	if user.Password != "" {
		err = user.HashPassword(user.Password)
	}
	return
}
