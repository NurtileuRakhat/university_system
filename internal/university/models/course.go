package models

import (
	"gorm.io/gorm"
	"time"
)

type Course struct {
	gorm.Model
	CourseName    string    `json:"course_name" gorm:"column:course_name;not null" binding:"required"`
	CourseCode    string    `json:"course_code" gorm:"column:course_code;not null" binding:"required"`
	CourseFaculty string    `json:"course_faculty" gorm:"column:course_faculty;not null" binding:"required"`
	Students      []Student `gorm:"many2many:student_courses;"`
	Teachers      []Teacher `gorm:"many2many:teacher_courses;"`
	CourseTime    time.Time `json:"course_time" gorm:"column:course_time;default:CURRENT_TIMESTAMP"`
}
