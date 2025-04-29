package models

type User struct {
	ID        string  `json:"id" db:"id"`
	Username  string  `json:"username" db:"username"`
	Password  string  `json:"password,omitempty" db:"password"`
	Firstname string  `json:"firstname" db:"firstname"`
	Lastname  string  `json:"lastname" db:"lastname"`
	Email     string  `json:"email" db:"email"`
	Role      string  `json:"role" db:"role"`
	Birthdate *string `json:"birthdate,omitempty" db:"birthdate"`
	CreatedAt string  `json:"created_at" db:"created_at"`
	UpdatedAt string  `json:"updated_at" db:"updated_at"`
	DeletedAt *string `json:"deleted_at,omitempty" db:"deleted_at"`
}
