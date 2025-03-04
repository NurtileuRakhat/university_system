package main

import (
	"fmt"
	"net/http"
	"os"
	_ "university_system/docs"
	"university_system/internal/routes"
	"university_system/pkg/databases"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
	logrus.SetFormatter(new(logrus.JSONFormatter))
	databases.Connect()
	databases.Migrate()
	router := gin.Default()
	routes.RegisterUserRoutes(router)
	logrus.Println(fmt.Sprintf("Listening on port %s", os.Getenv("PORT")))
	logrus.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router))
}
