package controller

import (
	"log"
	"net/http"
	"strconv"
	"university_system/internal/domain/models"
	"university_system/internal/university/services"

	"github.com/gin-gonic/gin"
)

type CourseController struct {
	courseService services.CourseService
}

func NewCourseController(service services.CourseService) *CourseController {
	return &CourseController{courseService: service}
}

// CreateCourse создает новый курс.
// @Summary Создать курс
// @Description Создает новый курс в системе
// @Tags courses
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer токен"
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
	// teacher_id может быть пустым (null) или uint, не забываем обработать
	createdCourse, err := c.courseService.CreateCourse(ctx.Request.Context(), &course)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create course: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdCourse)
}

// GetAllCourses получает список всех курсов.
// @Summary Получить все курсы
// @Description Возвращает список всех курсов в системе
// @Tags courses
// @Produce json
// @Param Authorization header string true "Bearer токен"
// @Success 200 {array} models.Course
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /courses [get]
// @Security BearerAuth
func (c *CourseController) GetAllCourses(ctx *gin.Context) {
	courses, err := c.courseService.GetAllCourses(ctx.Request.Context())
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
// @Param Authorization header string true "Bearer токен"
// @Param id path int true "ID курса"
// @Success 200 {object} models.Course
// @Failure 400 {object} gin.H "Неверный ID"
// @Failure 404 {object} gin.H "Курс не найден"
// @Router /courses/{id} [get]
// @Security BearerAuth
func (c *CourseController) GetCourseByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	// Проверка корректности ID опциональна, так как мы теперь работаем со строками
	// Но оставляем для совместимости с существующим кодом
	if _, err := strconv.ParseUint(idParam, 10, 64); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	course, err := c.courseService.GetCourseByID(ctx.Request.Context(), idParam)
	if err != nil {
		log.Println("Error fetching course:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch course"})
		return
	}

	if course == nil {
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
// @Param Authorization header string true "Bearer токен"
// @Param id path int true "ID курса"
// @Param course body models.Course true "Новые данные курса"
// @Success 200 {object} models.Course
// @Failure 400 {object} gin.H "Ошибка валидации"
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /courses/{id} [put]
// @Security BearerAuth
func (c *CourseController) UpdateCourse(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	var course models.Course
	if err := ctx.ShouldBindJSON(&course); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course.ID = string(id)
	updatedCourse, err := c.courseService.UpdateCourse(ctx.Request.Context(), course)
	if err != nil {
		log.Println("Error updating course:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update course"})
		return
	}

	ctx.JSON(http.StatusOK, updatedCourse)
}

// DeleteCourse удаляет курс по ID.
// @Summary Удалить курс
// @Description Удаляет курс из системы по его уникальному идентификатору
// @Tags courses
// @Param Authorization header string true "Bearer токен"
// @Param id path int true "ID курса"
// @Success 200 {object} gin.H "Курс удален"
// @Failure 400 {object} gin.H "Неверный ID"
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /courses/{id} [delete]
// @Security BearerAuth
func (c *CourseController) DeleteCourse(ctx *gin.Context) {
	idParam := ctx.Param("id")
	// Проверка корректности ID
	if _, err := strconv.ParseUint(idParam, 10, 64); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	err := c.courseService.DeleteCourse(ctx.Request.Context(), idParam)
	if err != nil {
		log.Println("Error deleting course:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete course"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}

// GetCourseStudents получает список студентов, записанных на курс.
// @Summary Получить студентов курса
// @Description Возвращает список студентов, записанных на данный курс
// @Tags courses
// @Produce json
// @Param Authorization header string true "Bearer токен"
// @Param id path int true "ID курса"
// @Success 200 {array} models.Student
// @Failure 400 {object} gin.H "Неверный ID"
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /courses/{id}/students [get]
// @Security BearerAuth
func (c *CourseController) GetCourseStudents(ctx *gin.Context) {
	idParam := ctx.Param("id")
	// Проверка корректности ID
	if _, err := strconv.ParseUint(idParam, 10, 64); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	students, err := c.courseService.GetCourseStudents(ctx.Request.Context(), idParam)
	if err != nil {
		log.Println("Error fetching course students:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch students"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"students": students})
}

// GetCourseTeachers получает список преподавателей курса.
// @Summary Получить преподавателей курса
// @Description Возвращает список преподавателей, ведущих данный курс
// @Tags courses
// @Produce json
// @Param Authorization header string true "Bearer токен"
// @Param id path int true "ID курса"
// @Success 200 {array} models.Teacher
// @Failure 400 {object} gin.H "Неверный ID"
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /courses/{id}/teachers [get]
// @Security BearerAuth
func (c *CourseController) GetCourseTeachers(ctx *gin.Context) {
	idParam := ctx.Param("id")
	// Проверка корректности ID
	if _, err := strconv.ParseUint(idParam, 10, 64); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	teachers, err := c.courseService.GetCourseTeachers(ctx.Request.Context(), idParam)
	if err != nil {
		log.Println("Error fetching course teachers:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch teachers"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"teachers": teachers})
}
