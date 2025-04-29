package models

import "errors"

type Mark struct {
	ID                uint    `json:"id" db:"id"`
	StudentID         uint    `json:"student_id" db:"student_id"`
	CourseID          uint    `json:"course_id" db:"course_id"`
	FirstAttestation  float64 `json:"first_attestation" db:"first_attestation"`
	SecondAttestation float64 `json:"second_attestation" db:"second_attestation"`
	FinalMark         float64 `json:"final_mark" db:"final_mark"`
	CreatedAt         string  `json:"created_at" db:"created_at"`
	UpdatedAt         string  `json:"updated_at" db:"updated_at"`
}

// Определение ошибок
var (
	ErrInvalidMarkType = errors.New("invalid mark type")
)
