package databases

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"university_system/internal/university/models"
)

var Instance *gorm.DB
var err error

func Connect() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		logrus.Fatalf("Missing required environment variables: DB_HOST=%s DB_PORT=%s DB_USER=%s DB_NAME=%s", dbHost, dbPort, dbUser, dbName)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	Instance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("Fatal connection to database: %v", err)
	}
	logrus.Println("Connected to Database")
}

func Migrate() {
	if err := Instance.AutoMigrate(&models.User{}, &models.Admin{}, &models.Course{}, &models.Student{}, &models.Teacher{}, &models.Manager{}, &models.CourseMark{}); err != nil {
		logrus.Fatalf("Fatal connection to database %s \n", err)
	}
	logrus.Println("Database Migration Completed...")
}
