package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"university_system/internal/domain/models"
	"university_system/internal/domain/repository"
)

type CourseMarkController struct {
	markRepo repository.GradeRepository
}

func NewCourseMarkController(repo repository.GradeRepository) *CourseMarkController {
	return &CourseMarkController{markRepo: repo}
}

func (c *CourseMarkController) addAttestationMark(ctx *gin.Context, markType string) {
	studentID := ctx.Param("student_id")
	courseID := ctx.Param("course_id")
	teacherID := ctx.Param("id")

	// Проверка, является ли преподаватель назначенным на данный курс
	isTeacher, err := c.markRepo.IsTeacherOfCourse(ctx.Request.Context(), teacherID, courseID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при проверке преподавателя"})
		return
	}
	if !isTeacher {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Преподаватель не назначен на данный курс"})
		return
	}

	var markValue float64
	if err := ctx.ShouldBindJSON(&markValue); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Некорректное значение оценки"})
		return
	}

	// Создание объекта Mark
	sid, err := strconv.ParseUint(studentID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный student_id"})
		return
	}
	cid, err := strconv.ParseUint(courseID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный course_id"})
		return
	}
	mark := &models.Mark{
		StudentID: uint(sid),
		CourseID:  uint(cid),
	}

	// Установка соответствующего значения оценки в зависимости от типа
	switch markType {
	case "first_attestation":
		mark.FirstAttestation = markValue
	case "second_attestation":
		mark.SecondAttestation = markValue
	case "final":
		mark.FinalMark = markValue
	}

	// Добавление оценки
	if err := c.markRepo.AddMark(ctx.Request.Context(), mark, markType); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось добавить оценку", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Оценка успешно добавлена"})
}

// AddFirstAttestation godoc
// @Summary Добавить первую аттестацию
// @Description Добавляет первую аттестацию студенту по курсу
// @Tags teachers
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer токен"
// @Param id path string true "ID учителя"
// @Param student_id path string true "ID студента"
// @Param course_id path string true "ID курса"
// @Param mark body float64 true "Оценка"
// @Success 201 {object} gin.H "Оценка добавлена"
// @Failure 400 {object} gin.H "Ошибка ввода"
// @Failure 403 {object} gin.H "Запрещено"
// @Router /teachers/{id}/courses/{course_id}/students/{student_id}/PutFirstAtt [post]
// @Security BearerAuth
func (c *CourseMarkController) AddFirstAttestation(ctx *gin.Context) {
	c.addAttestationMark(ctx, "first_attestation")
}

// AddSecondAttestation godoc
// @Summary Добавить вторую аттестацию
// @Description Добавляет вторую аттестацию студенту по курсу
// @Tags teachers
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer токен"
// @Param id path string true "ID учителя"
// @Param student_id path string true "ID студента"
// @Param course_id path string true "ID курса"
// @Param mark body float64 true "Оценка"
// @Success 201 {object} gin.H "Оценка добавлена"
// @Failure 400 {object} gin.H "Ошибка ввода"
// @Failure 403 {object} gin.H "Запрещено"
// @Router /teachers/{id}/courses/{course_id}/students/{student_id}/PutSecondAtt [post]
// @Security BearerAuth
func (c *CourseMarkController) AddSecondAttestation(ctx *gin.Context) {
	c.addAttestationMark(ctx, "second_attestation")
}

// AddFinalExamMark godoc
// @Summary Добавить итоговую оценку
// @Description Добавляет итоговую оценку студенту по курсу
// @Tags teachers
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer токен"
// @Param id path string true "ID учителя"
// @Param student_id path string true "ID студента"
// @Param course_id path string true "ID курса"
// @Param mark body float64 true "Оценка"
// @Success 201 {object} gin.H "Оценка добавлена"
// @Failure 400 {object} gin.H "Ошибка ввода"
// @Failure 403 {object} gin.H "Запрещено"
// @Router /teachers/{id}/courses/{course_id}/students/{student_id}/PutFinalMark [post]
// @Security BearerAuth
func (c *CourseMarkController) AddFinalExamMark(ctx *gin.Context) {
	c.addAttestationMark(ctx, "final")
}

// GetStudentMarks godoc
// @Summary Получить оценки студента
// @Description Возвращает все оценки студента по всем курсам.
// @Tags marks
// @Produce json
// @Param Authorization header string true "Bearer токен"
// @Param student_id path string true "ID студента"
// @Success 200 {object} models.Mark
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /marks/student/{student_id} [get]
// @Security BearerAuth
func (c *CourseMarkController) GetStudentMarks(ctx *gin.Context) {
	studentID := ctx.Param("student_id")

	marks, err := c.markRepo.GetStudentMarks(ctx.Request.Context(), studentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить оценки"})
		return
	}

	ctx.JSON(http.StatusOK, marks)
}

// GetCourseMarks godoc
// @Summary Получить оценки по курсу
// @Description Возвращает все оценки студентов по заданному курсу.
// @Tags marks
// @Produce json
// @Param Authorization header string true "Bearer токен"
// @Param course_id path string true "ID курса"
// @Success 200 {object} models.Mark
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /marks/course/{course_id} [get]
// @Security BearerAuth
func (c *CourseMarkController) GetCourseMarks(ctx *gin.Context) {
	courseID := ctx.Param("course_id")

	marks, err := c.markRepo.GetCourseMarks(ctx.Request.Context(), courseID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить оценки", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, marks)
}
