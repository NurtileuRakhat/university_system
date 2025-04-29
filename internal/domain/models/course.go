package models

type Course struct {
	ID          string  `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Code        string  `json:"code" db:"code"`
	Description string  `json:"description" db:"description"`
	TeacherID   *uint   `json:"teacher_id,omitempty" db:"teacher_id"`
	Credits     int     `json:"credits" db:"credits"`
	CreatedAt   string  `json:"created_at" db:"created_at"`
	UpdatedAt   string  `json:"updated_at" db:"updated_at"`
	DeletedAt   *string `json:"deleted_at,omitempty" db:"deleted_at"`
}
