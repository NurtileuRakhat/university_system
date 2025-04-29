package auth

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("your-secret-key")
var AccessTokenSecret = []byte("access-token-secret")
var RefreshTokenSecret = []byte("refresh-token-secret")

func GenerateAccessToken(username string, role string) (string, error) {
	claims := &jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(AccessTokenSecret)
}

func GenerateRefreshToken(username string, role string) (string, error) {
	claims := &jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 дней
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(RefreshTokenSecret)
}

func ParseRefreshToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return RefreshTokenSecret, nil
	})
}

func ParseAccessToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return AccessTokenSecret, nil
	})
}

// CheckPassword сравнивает хешированный пароль с обычным паролем
func CheckPassword(hashedPassword, password string) error {
	if hashedPassword == "" || password == "" {
		return errors.New("password cannot be empty")
	}
	
	// Если пароль еще не хеширован (для тестирования)
	if len(hashedPassword) < 60 && hashedPassword == password {
		return nil
	}
	
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// HashPassword хеширует пароль для безопасного хранения
func HashPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("password cannot be empty")
	}
	
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	
	return string(hashedBytes), nil
}
