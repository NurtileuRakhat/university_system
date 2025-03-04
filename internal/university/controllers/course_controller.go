package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"university_system/internal/university/models"
	"university_system/internal/university/repository"

	"github.com/gin-gonic/gin"
)

type CourseController struct {
	repo repository.CourseRepository
}

func NewCourseController(repo repository.CourseRepository) *CourseController {
	return &CourseController{repo: repo}
}

// CreateCourse создает новый курс.
// @Summary Создать курс
// @Description Создает новый курс в системе
// @Tags courses
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param course body models.Course true "Данные курса"
// @Success 201 {object} models.Course
// @Failure 400 {object} gin.H "Ошибка валидации"
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /courses [post]
// @Security BearerAuth
func (c *CourseController) CreateCourse(ctx *gin.Context) {
	var course models.Course
	if err := ctx.ShouldBindJSON(&course); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdUser, err := c.repo.CreateCourse(course)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create course: " + err.Error()})
		return
	}
	fmt.Printf("Saving Course: %+v\n", course)

	ctx.JSON(http.StatusCreated, createdUser)
}

// GetAllCourses получает список всех курсов.
// @Summary Получить все курсы
// @Description Возвращает список всех курсов в системе
// @Tags courses
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {array} models.Course
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /courses [get]
// @Security BearerAuth
func (c *CourseController) GetAllCourses(ctx *gin.Context) {
	courses, err := c.repo.GetAllCourses()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		return
	}
	ctx.JSON(http.StatusOK, courses)
}

// GetCourseByID получает курс по ID.
// @Summary Получить курс по ID
// @Description Возвращает курс по его уникальному идентификатору
// @Tags courses
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path int true "ID курса"
// @Success 200 {object} models.Course
// @Failure 400 {object} gin.H "Неверный ID"
// @Failure 404 {object} gin.H "Курс не найден"
// @Router /courses/{id} [get]
// @Security BearerAuth
func (c *CourseController) GetCourseByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	course, err := c.repo.GetCourseById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	ctx.JSON(http.StatusOK, course)
}

// UpdateCourse обновляет информацию о курсе.
// @Summary Обновить курс
// @Description Обновляет информацию о существующем курсе
// @Tags courses
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path int true "ID курса"
// @Param course body models.Course true "Новые данные курса"
// @Success 200 {object} models.Course
// @Failure 400 {object} gin.H "Ошибка валидации"
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /courses/{id} [put]
// @Security BearerAuth
func (c *CourseController) UpdateCourse(ctx *gin.Context) {
	id := ctx.Param("id")
	var course models.Course
	if err := ctx.ShouldBindJSON(&course); err != nil {
		log.Println("Error binding JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	courseID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Println("Error converting id:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	course.ID = uint(courseID)

	updatedUser, err := c.repo.UpdateCourse(course)
	if err != nil {
		log.Println("Error updating course:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update course"})
		return
	}

	ctx.JSON(http.StatusOK, updatedUser)
}

// DeleteCourse удаляет курс по ID.
// @Summary Удалить курс
// @Description Удаляет курс из системы по его уникальному идентификатору
// @Tags courses
// @Param Authorization header string true "Bearer Token"
// @Param id path int true "ID курса"
// @Success 200 {object} gin.H "Курс удален"
// @Failure 400 {object} gin.H "Неверный ID"
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /courses/{id} [delete]
// @Security BearerAuth
func (c *CourseController) DeleteCourse(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	if err := c.repo.DeleteCourse(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete course"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Course deleted"})
}

// GetCourseStudents получает список студентов, записанных на курс.
// @Summary Получить студентов курса
// @Description Возвращает список студентов, записанных на данный курс
// @Tags courses
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path int true "ID курса"
// @Success 200 {array} models.Student
// @Failure 400 {object} gin.H "Неверный ID"
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /courses/{id}/students [get]
// @Security BearerAuth
func (c *CourseController) GetCourseStudents(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	students, err := c.repo.GetCourseStudents(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch students"})
		return
	}
	ctx.JSON(http.StatusOK, students)
}

// GetCourseTeachers получает список преподавателей курса.
// @Summary Получить преподавателей курса
// @Description Возвращает список преподавателей, ведущих данный курс
// @Tags courses
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path int true "ID курса"
// @Success 200 {array} models.Teacher
// @Failure 400 {object} gin.H "Неверный ID"
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /courses/{id}/teachers [get]
// @Security BearerAuth
func (c *CourseController) GetCourseTeachers(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	teachers, err := c.repo.GetCourseTeachers(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch teachers"})
		return
	}
	ctx.JSON(http.StatusOK, teachers)
}
