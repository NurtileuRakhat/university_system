package auth

import (
	"log"
	"net/http"
	"university_system/internal/university/models"
	"university_system/internal/university/repository"
	"university_system/pkg/databases"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Login
// @Summary Login
// @Description Войти в систему, получив access и refresh токены
// @Tags Authorization
// @Accept json
// @Produce json
// @Param input body models.LoginRequest true "Данные пользователя"
// @Success 200 {object} models.AuthResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /login [post]
func Login(c *gin.Context) {
	var loginData models.LoginRequest

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid data"})
		return
	}

	userRepo := repository.NewUserRepository(databases.Instance)
	user, err := userRepo.GetUserByUsername(loginData.Username)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "User not found"})
		return
	}

	if err := user.CheckPassword(loginData.Password); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	accessToken, err := GenerateAccessToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Could not generate access token"})
		return
	}

	refreshToken, err := GenerateRefreshToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Could not generate refresh token"})
		return
	}

	c.JSON(http.StatusOK, models.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

// Refresh
// @Summary Обновление access-токена
// @Description Получение нового access-токена с помощью refresh-токена
// @Tags Authorization
// @Accept json
// @Produce json
// @Param input body models.RefreshRequest true "Refresh-токен"
// @Success 200 {object} models.AccessTokenResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /refresh [post]
func Refresh(c *gin.Context) {
	var refreshData models.RefreshRequest

	if err := c.ShouldBindJSON(&refreshData); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid data"})
		return
	}

	token, err := ParseRefreshToken(refreshData.RefreshToken)
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "Invalid or expired refresh token"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "Could not parse claims"})
		return
	}

	username := claims["username"].(string)
	accessToken, err := GenerateAccessToken(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Could not generate access token"})
		return
	}

	c.JSON(http.StatusOK, models.AccessTokenResponse{
		AccessToken: accessToken,
	})
}
