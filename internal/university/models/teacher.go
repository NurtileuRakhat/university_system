package models

type Teacher struct {
	User    `gorm:"embedded"`
	Courses []Course `gorm:"many2many:teacher_courses;"`
}
