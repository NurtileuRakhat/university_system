package models

type Teacher struct {
	User
	Department string  `json:"department" db:"department"`
	Position   string  `json:"position" db:"position"`
	CreatedAt  string  `json:"created_at" db:"created_at"`
	UpdatedAt  string  `json:"updated_at" db:"updated_at"`
	DeletedAt  *string `json:"deleted_at,omitempty" db:"deleted_at"`
}
