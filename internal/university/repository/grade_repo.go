package repository

import (
	"errors"
	"gorm.io/gorm"
	"university_system/internal/university/models"
)

type CourseMarkRepository interface {
	IsTeacherOfCourse(teacherID int, courseID int) bool
	AddMark(studentID, courseID uint, score float64, markType string) error
	GetFinalMarkAndGrade(studentID, courseID uint) (float64, string, error)
	GetStudentMarks(studentID uint) (courseMarks []models.CourseMark, err error)
	GetCourseMarks(courseID uint) ([]models.CourseMark, error)
}

type CourseMarkRepositoryImpl struct {
	DB    *gorm.DB
	grade models.CourseMark
}

func NewCourseMarkRepository(db *gorm.DB) *CourseRepositoryImpl {
	return &CourseRepositoryImpl{db: db}
}

func (r *CourseRepositoryImpl) IsTeacherOfCourse(teacherID int, courseID int) bool {
	var counter int64
	r.db.Model(&models.Teacher{}).Joins("JOIN teacher_courses ON teachers.id = teacher_courses.teacher_id").
		Where("teacher_courses.course_id = ? AND teachers.id = ?", courseID, teacherID).
		Count(&counter)

	return counter > 0
}
func (r *CourseRepositoryImpl) AddMark(studentID, courseID uint, score float64, markType string) error {
	var mark models.CourseMark
	if err := r.db.Where("student_id = ? AND course_id = ?", studentID, courseID).First(&mark).Error; err != nil {
		return errors.New("запись курса не найдена")
	}

	// Валидация оценки в зависимости от типа
	switch markType {
	case "first_attestation":
		if score < 0 || score > 30 {
			return errors.New("оценка за первую аттестацию должна быть в диапазоне 0-30")
		}
		mark.FirstAttestation = score
	case "second_attestation":
		if score < 0 || score > 30 {
			return errors.New("оценка за вторую аттестацию должна быть в диапазоне 0-30")
		}
		mark.SecondAttestation = score
	case "final_exam":
		if score < 0 || score > 40 {
			return errors.New("оценка за финальный экзамен должна быть в диапазоне 0-40")
		}
		mark.FinalExamMark = score
	default:
		return errors.New("неизвестный тип оценки")
	}

	// Пересчитываем общий балл и оценку
	mark.Recalculate()

	// Обновляем запись в БД
	return r.db.Save(&mark).Error
}

func (r *CourseRepositoryImpl) GetFinalMarkAndGrade(studentID, courseID uint) (float64, string, error) {
	var mark models.CourseMark
	if err := r.db.Where("student_id = ? AND course_id = ?", studentID, courseID).First(&mark).Error; err != nil {
		return 0, "", errors.New("запись курса не найдена")
	}
	finalMark, grade := mark.GetFinalMarkAndGrade()
	return finalMark, grade, nil
}

func (r *CourseRepositoryImpl) GetStudentMarks(studentID uint) ([]models.CourseMark, error) {
	var marks []models.CourseMark
	err := r.db.Where("student_id = ?", studentID).
		Preload("Course").
		Preload("Student").
		Preload("Teacher").
		Find(&marks).Error
	return marks, err
}

func (r *CourseRepositoryImpl) GetCourseMarks(courseID uint) ([]models.CourseMark, error) {
	var marks []models.CourseMark
	err := r.db.
		Where("course_id = ?", courseID).
		Preload("Student").
		Preload("Course").
		Preload("Teacher").
		Find(&marks).Error
	return marks, err
}
