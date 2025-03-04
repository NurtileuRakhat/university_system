package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	_ "university_system/internal/university/models"
	"university_system/internal/university/repository"
)

type CourseMarkController struct {
	markRepo repository.CourseMarkRepository
}

func NewCourseMarkController(repo repository.CourseMarkRepository) *CourseMarkController {
	return &CourseMarkController{markRepo: repo}
}

// addAttestationMark добавляет оценку студенту за определенный этап аттестации.
// @Summary Добавить оценку студенту
// @Description Добавляет оценку студенту по курсу, проверяя, является ли учитель преподавателем данного курса.
// @Tags marks
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param student_id path int true "ID студента"
// @Param course_id path int true "ID курса"
// @Param id path int true "ID учителя"
// @Param mark body float64 true "Оценка"
// @Success 201 {object} gin.H "Оценка добавлена"
// @Failure 400 {object} gin.H "Ошибка ввода"
// @Failure 403 {object} gin.H "Запрещено"
// @Router /marks/{id}/{student_id}/{course_id} [post]
// @Security BearerAuth
func (c *CourseMarkController) addAttestationMark(ctx *gin.Context, markType string) {
	studentID, _ := strconv.Atoi(ctx.Param("student_id"))
	courseID, _ := strconv.Atoi(ctx.Param("course_id"))
	teacherID, _ := strconv.Atoi(ctx.Param("id"))

	if !c.markRepo.IsTeacherOfCourse(teacherID, courseID) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Учитель не ведет этот курс"})
		return
	}

	var input struct {
		Mark float64 `json:"mark" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	err := c.markRepo.AddMark(uint(studentID), uint(courseID), input.Mark, markType)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Оценка добавлена"})
}

// @Security BearerAuth
func (c *CourseMarkController) AddFirstAttestation(ctx *gin.Context) {
	c.addAttestationMark(ctx, "first_attestation")
}

// @Security BearerAuth
func (c *CourseMarkController) AddSecondAttestation(ctx *gin.Context) {
	c.addAttestationMark(ctx, "second_attestation")
}

// @Security BearerAuth
func (c *CourseMarkController) AddFinalExamMark(ctx *gin.Context) {
	c.addAttestationMark(ctx, "final_exam")
}

// GetStudentMarks получает все оценки студента.
// @Summary Получить оценки студента
// @Description Возвращает все оценки студента по всем курсам.
// @Tags marks
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param student_id path int true "ID студента"
// @Success 200 {object} models.CourseMark
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /marks/student/{student_id} [get]
// @Security BearerAuth
func (c *CourseMarkController) GetStudentMarks(ctx *gin.Context) {
	studentID, _ := strconv.Atoi(ctx.Param("student_id"))

	marks, err := c.markRepo.GetStudentMarks(uint(studentID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения оценок"})
		return
	}

	ctx.JSON(http.StatusOK, marks)
}

// GetCourseMarks получает все оценки по конкретному курсу.
// @Summary Получить оценки по курсу
// @Description Возвращает все оценки студентов по заданному курсу.
// @Tags marks
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param course_id path int true "ID курса"
// @Success 200 {object} models.CourseMark
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /marks/course/{course_id} [get]
// @Security BearerAuth
func (c *CourseMarkController) GetCourseMarks(ctx *gin.Context) {
	courseID, _ := strconv.Atoi(ctx.Param("course_id"))

	marks, err := c.markRepo.GetCourseMarks(uint(courseID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения оценок"})
		return
	}

	ctx.JSON(http.StatusOK, marks)
}
