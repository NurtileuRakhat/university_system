package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"university_system/internal/domain/models"
	"university_system/internal/university/services"
)

type ManagerController struct {
	managerService services.ManagerService
}

func NewManagerController(managerService services.ManagerService) *ManagerController {
	return &ManagerController{managerService: managerService}
}

// GetManagers godoc
// @Summary Получить всех менеджеров
// @Description Возвращает список всех менеджеров из базы данных.
// @Tags managers
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer токен"
// @Success 200 {array} models.Manager "Список менеджеров"
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /managers [get]
func (mc *ManagerController) GetManagers(c *gin.Context) {
	managers, err := mc.managerService.GetManagers(c.Request.Context())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch managers"})
		return
	}
	c.JSON(http.StatusOK, managers)
}

// GetManagerById godoc
// @Summary Получить менеджера по ID
// @Description Возвращает данные менеджера по его идентификатору.
// @Tags managers
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID менеджера"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} models.Manager "Данные менеджера"
// @Failure 400 {object} gin.H "Некорректный ID"
// @Failure 404 {object} gin.H "Менеджер не найден"
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /managers/{id} [get]
func (mc *ManagerController) GetManagerById(c *gin.Context) {
	id := c.Param("id")
	manager, err := mc.managerService.GetManagerById(c.Request.Context(), id)
	if err != nil {
		log.Println("Error fetching manager:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch manager"})
		return
	}
	if manager == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Manager not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"manager": manager})
}

// CreateManager godoc
// @Summary Создать менеджера
// @Description Добавляет нового менеджера в систему.
// @Tags managers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer токен"
// @Param manager body models.Manager true "Данные нового менеджера"
// @Success 201 {object} models.Manager "Созданный менеджер"
// @Failure 400 {object} gin.H "Некорректные данные"
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /managers [post]
func (mc *ManagerController) CreateManager(c *gin.Context) {
	var manager models.Manager
	if err := c.ShouldBindJSON(&manager); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные: " + err.Error()})
		return
	}
	userID, err := mc.managerService.CreateUserWithRole(c.Request.Context(), manager.User, "manager")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания пользователя: " + err.Error()})
		return
	}
	manager.ID = userID
	manager.User.ID = userID
	createdManager, err := mc.managerService.CreateManager(c.Request.Context(), &manager)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания профиля менеджера: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdManager)
}

// UpdateManager godoc
// @Summary Обновить менеджера
// @Description Изменяет данные менеджера.
// @Tags managers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param manager body models.Manager true "Обновленные данные менеджера"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} models.Manager "Обновленный менеджер"
// @Failure 400 {object} gin.H "Некорректные данные"
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /managers [put]
func (mc *ManagerController) UpdateManager(c *gin.Context) {
	var manager models.Manager
	if err := c.ShouldBindJSON(&manager); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	updatedManager, err := mc.managerService.UpdateManager(c.Request.Context(), manager)
	if err != nil {
		log.Println("Error updating manager:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update manager"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"manager": updatedManager})
}

// DeleteManager godoc
// @Summary Удалить менеджера
// @Description Удаляет менеджера из системы по его идентификатору.
// @Tags managers
// @Security BearerAuth
// @Param id path string true "ID менеджера"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} gin.H "Успешное удаление"
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /managers/{id} [delete]
func (mc *ManagerController) DeleteManager(c *gin.Context) {
	id := c.Param("id")
	err := mc.managerService.DeleteManager(c.Request.Context(), id)
	if err != nil {
		log.Println("Error deleting manager:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete manager"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Manager deleted successfully"})
}

// AssignTeacherToCourse godoc
// @Summary Назначить преподавателя на курс
// @Description Привязывает преподавателя к курсу.
// @Tags managers
// @Security BearerAuth
// @Param teacher_id path string true "ID преподавателя"
// @Param Authorization header string true "Bearer токен"
// @Param course_id path string true "ID курса"
// @Success 200 {object} gin.H "Успешное назначение"
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /managers/assign/{teacher_id}/{course_id} [post]
func (mc *ManagerController) AssignTeacherToCourse(c *gin.Context) {
	teacherID := c.Param("teacher_id")
	courseID := c.Param("course_id")

	c.Set("teacher_id", teacherID)
	c.Set("course_id", courseID)

	log.Printf("Assigning teacher %s to course %s", teacherID, courseID)

	if err := mc.managerService.AssignTeacherToCourse(c.Request.Context(), teacherID, courseID); err != nil {
		log.Println("Error assigning teacher to course:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to assign teacher to course",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Teacher assigned to course successfully"})
}
