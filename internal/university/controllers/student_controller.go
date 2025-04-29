package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"university_system/internal/domain/models"
	"university_system/internal/university/services"
)

type StudentController struct {
	studentService services.StudentService
}

func NewStudentController(service services.StudentService) *StudentController {
	return &StudentController{studentService: service}
}

// GetStudents godoc
// @Summary Получить список всех студентов
// @Description Возвращает список всех студентов в системе
// @Tags students
// @Security BearerAuth
// @Param Authorization header string true "Bearer токен"
// @Accept json
// @Produce json
// @Success 200 {array} models.Student
// @Failure 401 {object} map[string]string "Неавторизованный доступ"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /students [get]
func (sc *StudentController) GetStudents(c *gin.Context) {
	students, err := sc.studentService.GetStudents(c.Request.Context())
	if err != nil {
		log.Println("Error fetching students:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch students"})
		return
	}
	c.JSON(http.StatusOK, students)
}

// GetStudentById godoc
// @Summary Получить информацию о студенте
// @Description Возвращает данные студента по его ID
// @Tags students
// @Security BearerAuth
// @Param Authorization header string true "Bearer токен"
// @Param id path string true "ID студента"
// @Accept json
// @Produce json
// @Success 200 {object} models.Student
// @Failure 400 {object} map[string]string "Неверный запрос"
// @Failure 401 {object} map[string]string "Неавторизованный доступ"
// @Failure 404 {object} map[string]string "Студент не найден"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /students/{id} [get]
func (sc *StudentController) GetStudentById(ctx *gin.Context) {
	id := ctx.Param("id")
	student, err := sc.studentService.GetStudentById(ctx.Request.Context(), id)
	if err != nil {
		log.Println("Error fetching student:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch student"})
		return
	}
	if student == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, student)
}

// CreateStudent godoc
// @Summary Создать нового студента
// @Description Добавляет нового студента в систему
// @Tags students
// @Security BearerAuth
// @Param Authorization header string true "Bearer токен"
// @Param input body models.Student true "Данные нового студента"
// @Accept json
// @Produce json
// @Success 201 {object} models.Student
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /students [post]
func (sc *StudentController) CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные: " + err.Error()})
		return
	}
	userID, err := sc.studentService.CreateUserWithRole(c.Request.Context(), student.User, "student")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания пользователя: " + err.Error()})
		return
	}
	student.ID = userID
	student.User.ID = userID
	createdStudent, err := sc.studentService.CreateStudent(c.Request.Context(), &student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания профиля студента: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdStudent)
}

// UpdateStudent godoc
// @Summary Обновить данные студента
// @Description Обновляет информацию о студенте
// @Tags students
// @Security BearerAuth
// @Param Authorization header string true "Bearer токен"
// @Param id path string true "ID студента"
// @Param input body models.Student true "Обновленные данные студента"
// @Accept json
// @Produce json
// @Success 200 {object} models.Student
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /students/{id} [put]
func (sc *StudentController) UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	student.ID = id

	updatedUser, err := sc.studentService.UpdateStudent(c.Request.Context(), student)
	if err != nil {
		log.Println("Error updating student:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update student"})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// DeleteStudent godoc
// @Summary Удалить студента
// @Description Удаляет студента по его ID
// @Tags students
// @Security BearerAuth
// @Param Authorization header string true "Bearer токен"
// @Param id path string true "ID студента"
// @Success 204
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /students/{id} [delete]
func (sc *StudentController) DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	if err := sc.studentService.DeleteStudent(c.Request.Context(), id); err != nil {
		log.Println("Error deleting student:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete student"})
		return
	}
	c.Status(http.StatusNoContent)
}

// EnrollStudentToCourse godoc
// @Summary Записать студента на курс
// @Description Записывает студента на указанный курс
// @Tags students
// @Security BearerAuth
// @Param Authorization header string true "Bearer токен"
// @Param student_id path string true "ID студента"
// @Param course_id path string true "ID курса"
// @Success 200 {object} map[string]string
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /students/{student_id}/courses/{course_id} [post]
func (sc *StudentController) EnrollStudentToCourse(ctx *gin.Context) {
	studentID := ctx.Param("student_id")
	courseID := ctx.Param("course_id")

	ctx.Set("student_id", studentID)
	ctx.Set("course_id", courseID)

	if err := sc.studentService.EnrollStudentToCourse(ctx.Request.Context(), studentID, courseID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка записи на курс", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Студент успешно записан на курс"})
}

// GetStudentCourses godoc
// @Summary Получить курсы студента
// @Description Возвращает список курсов, на которые записан студент
// @Tags students
// @Security BearerAuth
// @Param Authorization header string true "Bearer токен"
// @Param id path string true "ID студента"
// @Accept json
// @Produce json
// @Success 200 {array} models.Course
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /students/{id}/courses [get]
func (sc *StudentController) GetStudentCourses(ctx *gin.Context) {
	studentID := ctx.Param("id")

	courses, err := sc.studentService.GetStudentCourses(ctx.Request.Context(), studentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch student courses"})
		return
	}

	ctx.JSON(http.StatusOK, courses)
}
