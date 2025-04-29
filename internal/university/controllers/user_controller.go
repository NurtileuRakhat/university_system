package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"university_system/internal/domain/models"
	"university_system/internal/university/services"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{UserService: userService}
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
	ctx := c.Request.Context()
	users, err := uc.UserService.GetUsers(ctx)
	if err != nil {
		log.Println("Error fetching users:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// CreateUser
// @Summary Создание пользователя
// @Description Создаёт нового пользователя в системе
// @Tags users
// @Accept  json
// @Produce  json
// @Param input body models.User true "Данные пользователя"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/users [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind JSON"})
		return
	}
	ctx := c.Request.Context()
	createdUser, err := uc.UserService.CreateUser(ctx, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, createdUser)
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
	ctx := c.Request.Context()
	user, err := uc.UserService.GetUserById(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind JSON"})
		return
	}
	user.ID = id
	ctx := c.Request.Context()
	updatedUser, err := uc.UserService.UpdateUser(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	ctx := c.Request.Context()
	err := uc.UserService.DeleteUser(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
