package models

import (
	"gorm.io/gorm"
)

type CourseMark struct {
	gorm.Model
	StudentID         uint    `json:"student_id"`
	CourseID          uint    `json:"course_id"`
	TeacherID         uint    `json:"teacher_id"`
	FirstAttestation  float64 `json:"first_attestation"`
	SecondAttestation float64 `json:"second_attestation"`
	FinalExamMark     float64 `json:"final_exam_mark"`
	Mark              float64 `json:"mark"`
	Grade             string  `json:"grade"`
	Student           Student `gorm:"foreignKey:StudentID"`
	Course            Course  `gorm:"foreignKey:CourseID"`
	Teacher           Teacher `gorm:"foreignKey:TeacherID"`
}

func (cm *CourseMark) Recalculate() {
	cm.Mark = cm.FirstAttestation + cm.SecondAttestation + cm.FinalExamMark
	cm.Grade = cm.ConvertToGrade()
}

func (cm *CourseMark) GetFinalMarkAndGrade() (float64, string) {
	return cm.Mark, cm.Grade
}

func (cm *CourseMark) ConvertToGrade() string {
	switch {
	case cm.Mark >= 90:
		return "A"
	case cm.Mark >= 85:
		return "A-"
	case cm.Mark >= 80:
		return "B+"
	case cm.Mark >= 75:
		return "B"
	case cm.Mark >= 70:
		return "B-"
	case cm.Mark >= 65:
		return "C+"
	case cm.Mark >= 60:
		return "C"
	case cm.Mark >= 55:
		return "C-"
	case cm.Mark >= 50:
		return "D"
	default:
		return "F"
	}
}
