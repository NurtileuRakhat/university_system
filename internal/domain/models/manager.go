package models

type Manager struct {
	User
	Department string  `json:"department" db:"department"`
	CreatedAt  string  `json:"created_at" db:"created_at"`
	UpdatedAt  string  `json:"updated_at" db:"updated_at"`
	DeletedAt  *string `json:"deleted_at,omitempty" db:"deleted_at"`
}
