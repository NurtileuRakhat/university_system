package auth

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

var enforcer *casbin.Enforcer

func InitCasbin() error {
	modelPath := filepath.Join("config", "rbac_model.conf")
	adapter := filepath.Join("config", "rbac_policy.csv")

	var err error
	enforcer, err = casbin.NewEnforcer(modelPath, adapter)
	if err != nil {
		return err
	}

	// Загружаем политики
	if err := enforcer.LoadPolicy(); err != nil {
		return err
	}

	return nil
}

func CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем роль пользователя из контекста
		userRole, exists := c.Get("user_role")
		if !exists {
			log.Printf("Role not found in context")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Проверяем разрешения
		obj := c.Request.URL.Path
		act := c.Request.Method

		log.Printf("Casbin middleware - Role from context: %v (type: %T)", userRole, userRole)
		log.Printf("Checking permissions for role: %v, path: %s, method: %s", userRole, obj, act)

		// Преобразуем роль в строку, если это не строка
		roleStr, ok := userRole.(string)
		if !ok {
			log.Printf("Error: role is not a string, it is %T", userRole)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid role format"})
			c.Abort()
			return
		}

		ok, err := enforcer.Enforce(roleStr, obj, act)
		if err != nil {
			log.Printf("Error checking permissions: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking permissions"})
			c.Abort()
			return
		}

		if !ok {
			log.Printf("Access denied for role: %v, path: %s, method: %s", roleStr, obj, act)
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}

		log.Printf("Access granted for role: %v, path: %s, method: %s", roleStr, obj, act)
		c.Next()
	}
}

// AddPolicy добавляет новую политику
func AddPolicy(sub, obj, act string) error {
	_, err := enforcer.AddPolicy(sub, obj, act)
	return err
}

// RemovePolicy удаляет политику
func RemovePolicy(sub, obj, act string) error {
	_, err := enforcer.RemovePolicy(sub, obj, act)
	return err
}

// AddRoleForUser назначает роль пользователю
func AddRoleForUser(user, role string) error {
	_, err := enforcer.AddRoleForUser(user, role)
	return err
}

// RemoveRoleForUser удаляет роль у пользователя
func RemoveRoleForUser(user, role string) error {
	_, err := enforcer.DeleteRoleForUser(user, role)
	return err
}
