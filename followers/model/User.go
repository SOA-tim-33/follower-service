package model

import (
	"database-example/model/enums"
)

type User struct {
	Id       int    `gorm:"primary_key;auto_increment"`
	Username string `gorm:"unique"`
	Password string `gorm:"not null"`
	Role     enums.UserRole
	IsActive bool
	Email    string `gorm:"unique"`
	Token    string
}
