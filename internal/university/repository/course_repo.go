package repository

import (
	"gorm.io/gorm"
	"university_system/internal/university/models"
)

type CourseRepository interface {
	CreateCourse(course models.Course) (*models.Course, error)
	GetAllCourses() ([]models.Course, error)
	GetCourseById(id uint) (*models.Course, error)
	UpdateCourse(course models.Course) (*models.Course, error)
	DeleteCourse(id uint) error
	GetCourseStudents(courseId int) ([]models.Student, error)
	GetCourseTeachers(courseId int) ([]models.Teacher, error)
}

type CourseRepositoryImpl struct {
	db     *gorm.DB
	course models.Course
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &CourseRepositoryImpl{db: db}
}

func (r *CourseRepositoryImpl) CreateCourse(course models.Course) (*models.Course, error) {
	if err := r.db.Create(&course).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *CourseRepositoryImpl) GetAllCourses() ([]models.Course, error) {
	var courses []models.Course
	if err := r.db.Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *CourseRepositoryImpl) GetCourseById(id uint) (*models.Course, error) {
	var course models.Course
	if err := r.db.First(&course, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *CourseRepositoryImpl) UpdateCourse(course models.Course) (*models.Course, error) {
	if err := r.db.Save(&course).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *CourseRepositoryImpl) DeleteCourse(id uint) error {
	if err := r.db.Delete(&models.Course{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
func (r *CourseRepositoryImpl) GetCourseStudents(courseId int) ([]models.Student, error) {
	var course models.Course

	if err := r.db.Preload("Students").First(&course, "id = ?", courseId).Error; err != nil {
		return nil, err
	}
	return course.Students, nil
}

func (r *CourseRepositoryImpl) GetCourseTeachers(courseId int) ([]models.Teacher, error) {
	var course models.Course

	if err := r.db.Preload("Teachers").First(&course, "id = ?", courseId).Error; err != nil {
		return nil, err
	}
	return course.Teachers, nil
}
