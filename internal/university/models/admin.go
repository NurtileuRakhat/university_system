package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	User
	gorm.Model
}
