package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"university_system/internal/university/models"
	"university_system/internal/university/repository"
)

type UserController struct {
	UserRepo repository.UserRepository
}

func NewUserController(userRepo repository.UserRepository) *UserController {
	return &UserController{UserRepo: userRepo}
}

// GetUsers
// @Summary Get all users
// @Description Возвращает всех пользователей системы
// @Tags users
// @Security BearerAuth
// @Param Authorization header string true "Bearer токен"
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/users [get]
func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.UserRepo.GetUsers()
	if err != nil {
		log.Println("Error fetching users:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// Register
// @Summary Создание пользователя
// @Description Создаёт нового пользователя в системе
// @Tags Authorization
// @Accept  json
// @Produce  json
// @Param input body models.User true "Данные пользователя"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /register [post]
func (uc *UserController) Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind JSON"})
		return
	}

	createdUser, err := uc.UserRepo.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

// GetUserById
// @Summary get user by id
// @Description get user by id
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer токен"
// @Param id path string true "ID пользователя"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/users/{id} [get]
func (uc *UserController) GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := uc.UserRepo.GetUserById(id)
	if err != nil {
		log.Println("Error fetching user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch user"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser
// @Summary Обновление данных пользователя
// @Description Обновляет информацию о пользователе по его ID
// @Tags users
// @Security BearerAuth
// @Param Authorization header string true "Bearer токен"
// @Param id path string true "ID пользователя"
// @Param input body models.User true "Обновлённые данные пользователя"
// @Param Authorization header string true "Bearer токен"
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/users/{id} [put]
func (uc *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Println("Error converting id:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	user.ID = uint(userID)

	updatedUser, err := uc.UserRepo.UpdateUser(user)
	if err != nil {
		log.Println("Error updating user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update user"})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser
// @Summary Удаление пользователя
// @Description Удаляет пользователя по его ID
// @Tags users
// @Security BearerAuth
// @Param Authorization header string true "Bearer токен"
// @Param id path string true "ID пользователя"
// @Success 204 "Пользователь успешно удалён"
// @Failure 500 {object} models.ErrorResponse
// @Router /api/users/{id} [delete]
func (uc *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := uc.UserRepo.DeleteUser(id); err != nil {
		log.Println("Error deleting user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete user"})
		return
	}
	c.Status(http.StatusNoContent)
}
