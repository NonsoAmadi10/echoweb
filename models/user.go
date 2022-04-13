package models

import (
	"github.com/NonsoAmadi10/echoweb/utils"
	"gorm.io/gorm"
)

type User struct {
	Model
	Email    string `gorm:"uniqueIndex"`
	FullName string
	Password string
	Username string `gorm:"uniqueIndex;not_null"`
	Role     string
}

func (user User) String() string {
	return user.FullName
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	hashed, err := utils.HashPassword(user.Password)
	user.Password = hashed

	return
}
