package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"university_system/internal/domain/models"
	"university_system/internal/university/services"
)

type TeacherController struct {
	teacherService services.TeacherService
}

func NewTeacherController(teacherService services.TeacherService) *TeacherController {
	return &TeacherController{teacherService: teacherService}
}

// GetTeachers godoc
// @Summary Получить список всех преподавателей
// @Description Возвращает список всех преподавателей в системе
// @Tags teachers
// @Security BearerAuth
// @Param Authorization header string true "Bearer токен"
// @Accept json
// @Produce json
// @Success 200 {array} models.Teacher
// @Failure 500 {object} models.ErrorResponse
// @Router /teachers [get]
func (tc *TeacherController) GetTeachers(c *gin.Context) {
	teachers, err := tc.teacherService.GetTeachers(c.Request.Context())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch teachers"})
		return
	}
	c.JSON(http.StatusOK, teachers)
}

// GetTeacherByID godoc
// @Summary Получить информацию о преподавателе
// @Description Возвращает данные преподавателя по его ID
// @Tags teachers
// @Security BearerAuth
// @Param Authorization header string true "Bearer токен"
// @Param id path string true "ID преподавателя"
// @Accept json
// @Produce json
// @Success 200 {object} models.Teacher
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /teachers/{id} [get]
func (tc *TeacherController) GetTeacherByID(c *gin.Context) {
	id := c.Param("id")
	teacher, err := tc.teacherService.GetTeacherById(c.Request.Context(), id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch teacher"})
		return
	}
	if teacher == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Teacher not found"})
		return
	}
	c.JSON(http.StatusOK, teacher)
}

// CreateTeacher godoc
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
// @Router /teachers [post]
func (tc *TeacherController) CreateTeacher(c *gin.Context) {
	var teacher models.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные: " + err.Error()})
		return
	}
	// 1. Создать пользователя с ролью teacher
	userID, err := tc.teacherService.CreateUserWithRole(c.Request.Context(), teacher.User, "teacher")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания пользователя: " + err.Error()})
		return
	}
	teacher.ID = userID
	teacher.User.ID = userID
	// 2. Создать преподавателя
	createdTeacher, err := tc.teacherService.CreateTeacher(c.Request.Context(), &teacher)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания профиля преподавателя: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdTeacher)
}

// UpdateTeacher godoc
// @Summary Обновить данные преподавателя
// @Description Обновляет информацию о преподавателе
// @Tags teachers
// @Security BearerAuth
// @Param Authorization header string true "Bearer токен"
// @Param id path string true "ID преподавателя"
// @Param input body models.Teacher true "Обновленные данные преподавателя"
// @Accept json
// @Produce json
// @Success 200 {object} models.Teacher
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /teachers/{id} [put]
func (tc *TeacherController) UpdateTeacher(c *gin.Context) {
	id := c.Param("id")
	var teacher models.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	teacher.ID = id
	updatedTeacher, err := tc.teacherService.UpdateTeacher(c.Request.Context(), teacher)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update teacher"})
		return
	}
	c.JSON(http.StatusOK, updatedTeacher)
}

// DeleteTeacher godoc
// @Summary Удалить преподавателя
// @Description Удаляет преподавателя по его ID
// @Tags teachers
// @Security BearerAuth
// @Param Authorization header string true "Bearer токен"
// @Param id path string true "ID преподавателя"
// @Success 200 {object} map[string]string
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /teachers/{id} [delete]
func (tc *TeacherController) DeleteTeacher(c *gin.Context) {
	id := c.Param("id")
	if err := tc.teacherService.DeleteTeacher(c.Request.Context(), id); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete teacher"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Teacher deleted successfully"})
}

// GetTeacherCourses godoc
// @Summary Получить курсы преподавателя
// @Description Возвращает список курсов, которые ведет преподаватель
// @Tags teachers
// @Security BearerAuth
// @Param Authorization header string true "Bearer токен"
// @Param id path string true "ID преподавателя"
// @Accept json
// @Produce json
// @Success 200 {array} models.Course
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /teachers/{id}/courses [get]
func (tc *TeacherController) GetTeacherCourses(c *gin.Context) {
	id := c.Param("id")
	courses, err := tc.teacherService.GetTeacherCourses(c.Request.Context(), id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		return
	}
	c.JSON(http.StatusOK, courses)
}
