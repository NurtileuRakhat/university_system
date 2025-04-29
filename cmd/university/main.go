package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	_ "university_system/docs"
	"university_system/internal/routes"
	"university_system/pkg/config"
	"university_system/pkg/databases"
)

// @title University System
// @version 1.0
// @description System для управления пользователями, курсами и ролями
// @contact.name Rakhat Nurtileu
// @contact.url http://linkedin.com/in/nurtileu-rakhat-02aba3294/
// @contact.email n_rakhat@kbtu.kz
// @host localhost:8080
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg := config.LoadConfig()
	pgURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBName, cfg.DB.SSLMode)

	logrus.Infof("Connecting to DB: %s", pgURL)
	db, err := sqlx.Connect("postgres", pgURL)
	if err != nil {
		logrus.Errorf("Failed to connect to DB: %v", err)
		return
	}

	databases.Instance = db
	logrus.Info("Successfully connected to DB")

	ctx := context.Background()

	if err := databases.Migrate(ctx); err != nil {
		logrus.Errorf("Migration failed: %v", err)
		return
	}

	logrus.SetFormatter(new(logrus.JSONFormatter))
	router := gin.Default()
	routes.RegisterUserRoutes(router)
	logrus.Println(fmt.Sprintf("Listening on port %s", os.Getenv("PORT")))
	logrus.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router))
}
