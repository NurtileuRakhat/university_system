package models

type Student struct {
	User
	StudentYear int     `json:"student_year" db:"student_year"`
	Faculty     string  `json:"faculty" db:"faculty"`
	CreatedAt   string  `json:"created_at" db:"created_at"`
	UpdatedAt   string  `json:"updated_at" db:"updated_at"`
	DeletedAt   *string `json:"deleted_at,omitempty" db:"deleted_at"`
}
