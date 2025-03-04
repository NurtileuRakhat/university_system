package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"university_system/internal/university/models"
	"university_system/internal/university/repository"
)

type StudentController struct {
	studentRepo repository.StudentRepository
}

func NewStudentController(studentRepo repository.StudentRepository) *StudentController {
	return &StudentController{studentRepo: studentRepo}
}

// GetStudents
// @Summary Получить список всех студентов
// @Description Возвращает список всех студентов в системе
// @Tags students
// @Security BearerAuth
// @Param Authorization: Bearer header string true "Bearer токен"
// @Accept json
// @Produce json
// @Success 200 {array} models.Student
// @Failure 500 {object} models.ErrorResponse
// @Router /api/students [get]
func (sc *StudentController) GetStudents(c *gin.Context) {
	students, err := sc.studentRepo.GetStudents()
	if err != nil {
		log.Println("Error fetching students:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch students"})
		return
	}
	c.JSON(http.StatusOK, students)
}

// GetStudentById
// @Summary Получить информацию о студенте
// @Description Возвращает данные студента по его ID
// @Tags students
// @Security BearerAuth
// @Param Authorization header string true "Bearer токен"
// @Param id path string true "ID студента"
// @Accept json
// @Produce json
// @Success 200 {object} models.Student
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/students/{id} [get]
func (sc *StudentController) GetStudentById(ctx *gin.Context) {
	id := ctx.Param("id")
	student, err := sc.studentRepo.GetStudentsById(id)
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

// CreateStudent
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
// @Router /api/students [post]
func (sc *StudentController) CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind JSON" + err.Error()})
		return
	}

	createdUser, err := sc.studentRepo.CreateStudent(&student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

// UpdateStudent
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
// @Router /api/students/{id} [put]
func (sc *StudentController) UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	studentID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Println("Error converting id:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	student.ID = uint(studentID)

	updatedUser, err := sc.studentRepo.UpdateStudent(&student)
	if err != nil {
		log.Println("Error updating student:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update student"})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// DeleteStudent
// @Summary Удалить студента
// @Description Удаляет студента по его ID
// @Tags students
// @Security BearerAuth
// @Param Authorization header string true "Bearer токен"
// @Param id path string true "ID студента"
// @Success 204
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/students/{id} [delete]
func (sc *StudentController) DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	if err := sc.studentRepo.DeleteStudent(id); err != nil {
		log.Println("Error deleting student:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete student"})
		return
	}
	c.Status(http.StatusNoContent)
}

// EnrollStudentToCourse
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
// @Router /api/students/{student_id}/courses/{course_id} [post]
func (c *StudentController) EnrollStudentToCourse(ctx *gin.Context) {
	studentID := ctx.Param("student_id")
	courseID := ctx.Param("course_id")

	ctx.Set("student_id", studentID)
	ctx.Set("course_id", courseID)

	if err := c.studentRepo.EnrollStudentToCourse(studentID, courseID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка записи на курс", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Студент успешно записан на курс"})
}

// GetStudentCourses
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
// @Router /api/students/{id}/courses [get]
func (sc *StudentController) GetStudentCourses(ctx *gin.Context) {
	studentID := ctx.Param("id")

	courses, err := sc.studentRepo.GetStudentCourses(studentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения курсов студента", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, courses)
}
