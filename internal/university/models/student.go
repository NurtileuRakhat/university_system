package models

type Student struct {
	User        `gorm:"embedded"`
	StudentYear int      `json:"student_year" gorm:"column:student_year;not null"`
	Faculty     string   `json:"faculty" gorm:"column:faculty;not null"`
	Courses     []Course `gorm:"many2many:student_courses; "`
}
