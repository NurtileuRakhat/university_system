package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"university_system/internal/university/models"
	"university_system/internal/university/repository"
)

type TeacherController struct {
	teacherRepo repository.TeacherRepository
}

func NewTeacherController(teacherRepo repository.TeacherRepository) *TeacherController {
	return &TeacherController{teacherRepo: teacherRepo}
}

// GetTeachers
// @Summary Получить список всех преподавателей
// @Description Возвращает список всех преподавателей в системе
// @Tags teachers
// @Security BearerAuth
// @Param Authorization header string true "Bearer токен"
// @Accept json
// @Produce json
// @Success 200 {array} models.Teacher
// @Failure 500 {object} models.ErrorResponse
// @Router /api/teachers [get]
func (tc *TeacherController) GetTeachers(c *gin.Context) {
	teachers, err := tc.teacherRepo.GetTeachers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch teachers"})
		return
	}
	c.JSON(http.StatusOK, teachers)
}

// GetTeacherByID
// @Summary Получить информацию о преподавателе
// @Description Возвращает данные преподавателя по его ID
// @Tags teachers
// @Security BearerAuth
// @Param Authorization header string true "Bearer токен"
// @Param id path int true "ID преподавателя"
// @Accept json
// @Produce json
// @Success 200 {object} models.Teacher
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/teachers/{id} [get]
func (tc *TeacherController) GetTeacherByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid teacher ID"})
		return
	}

	teacher, err := tc.teacherRepo.GetTeacherByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch teacher"})
		return
	}
	if teacher == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Teacher not found"})
		return
	}
	c.JSON(http.StatusOK, teacher)
}

// CreateTeacher
// @Summary Создать нового преподавателя
// @Description Добавляет нового преподавателя в систему
// @Tags teachers
// @Security BearerAuth
// @Param Authorization header string true "Bearer токен"
// @Param input body models.Teacher true "Данные нового преподавателя"
// @Accept json
// @Produce json
// @Success 201 {object} models.Teacher
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/teachers [post]
func (tc *TeacherController) CreateTeacher(c *gin.Context) {
	var teacher models.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	createdTeacher, err := tc.teacherRepo.CreateTeacher(&teacher)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create teacher"})
		return
	}
	c.JSON(http.StatusCreated, createdTeacher)
}

// UpdateTeacher
// @Summary Обновить данные преподавателя
// @Description Обновляет информацию о преподавателе
// @Tags teachers
// @Security BearerAuth
// @Param Authorization header string true "Bearer токен"
// @Param input body models.Teacher true "Обновленные данные преподавателя"
// @Accept json
// @Produce json
// @Success 200 {object} models.Teacher
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/teachers [put]
func (tc *TeacherController) UpdateTeacher(c *gin.Context) {
	var teacher models.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	updatedTeacher, err := tc.teacherRepo.UpdateTeacher(&teacher)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update teacher"})
		return
	}
	c.JSON(http.StatusOK, updatedTeacher)
}

// DeleteTeacher
// @Summary Удалить преподавателя
// @Description Удаляет преподавателя по его ID
// @Tags teachers
// @Security BearerAuth
// @Param Authorization header string true "Bearer токен"
// @Param id path int true "ID преподавателя"
// @Success 200 {object} map[string]string
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/teachers/{id} [delete]
func (tc *TeacherController) DeleteTeacher(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid teacher ID"})
		return
	}

	if err := tc.teacherRepo.DeleteTeacher(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete teacher"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Teacher deleted successfully"})
}

// GetTeacherCourses
// @Summary Получить курсы преподавателя
// @Description Возвращает список курсов, которые ведет преподаватель
// @Tags teachers
// @Security BearerAuth
// @Param Authorization header string true "Bearer токен"
// @Param id path int true "ID преподавателя"
// @Accept json
// @Produce json
// @Success 200 {array} models.Course
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/teachers/{id}/courses [get]
func (tc *TeacherController) GetTeacherCourses(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid teacher ID"})
		return
	}

	courses, err := tc.teacherRepo.GetTeacherCourses(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		return
	}
	c.JSON(http.StatusOK, courses)
}
