package models

// LoginRequest представляет данные запроса для входа в систему
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RefreshRequest представляет данные запроса для обновления токена
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// AuthResponse представляет ответ с токенами авторизации
type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// AccessTokenResponse представляет ответ только с access токеном
type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

// ErrorResponse представляет ответ с сообщением об ошибке
type ErrorResponse struct {
	Message string `json:"message"`
}
