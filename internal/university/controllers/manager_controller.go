package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"university_system/internal/university/models"
	"university_system/internal/university/repository"
)

type ManagerController struct {
	managerRepo repository.ManagerRepository
}

func NewManagerController(managerRepo repository.ManagerRepository) *ManagerController {
	return &ManagerController{managerRepo: managerRepo}
}

// GetManagers получает список всех менеджеров.
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
	managers, err := mc.managerRepo.GetManagers()
	if err != nil {
		log.Println("Error fetching managers:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch managers"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"managers": managers})
}

// GetManagerById получает менеджера по ID.
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
	manager, err := mc.managerRepo.GetManagerById(id)
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

// CreateManager создает нового менеджера.
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	newManager, err := mc.managerRepo.CreateManager(manager)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create manager"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"manager": newManager})
}

// UpdateManager обновляет данные менеджера.
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	updatedManager, err := mc.managerRepo.UpdateManager(&manager)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update manager"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"manager": updatedManager})
}

// DeleteManager удаляет менеджера по ID.
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
	if err := mc.managerRepo.DeleteManager(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete manager"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Manager deleted successfully"})
}

// AssignTeacherToCourse назначает преподавателя на курс.
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

	if err := mc.managerRepo.AssignTeacherToCourse(teacherID, courseID); err != nil {
		log.Println("Error assigning teacher to course:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to assign teacher to course",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Teacher assigned to course successfully"})
}
