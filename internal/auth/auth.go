package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"university_system/internal/domain/models"
	infraRepo "university_system/internal/infrastructure/repository"
	"university_system/pkg/databases"
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

	userRepo := infraRepo.NewUserRepository(databases.Instance)
	user, err := userRepo.GetUserByUsername(c.Request.Context(), loginData.Username)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "User not found"})
		return
	}

	if err := CheckPassword(user.Password, loginData.Password); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	accessToken, err := GenerateAccessToken(user.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Could not generate access token"})
		return
	}

	refreshToken, err := GenerateRefreshToken(user.Username, user.Role)
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

	token, err := jwt.Parse(refreshData.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(RefreshTokenSecret), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "Invalid refresh token"})
		return
	}

	if !token.Valid {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "Invalid refresh token"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "Invalid token claims"})
		return
	}

	username, ok := claims["username"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "Invalid username claim"})
		return
	}

	accessToken, err := GenerateAccessToken(username, claims["role"].(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Could not generate access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}
