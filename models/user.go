package models

import (
	Common "github.com/NonsoAmadi10/echoweb/common"
	"github.com/NonsoAmadi10/echoweb/utils"
)

type User struct {
	Common.Base 
	Email string `gorm:"type:varchar(100);unique_index"`
	FullName     string
	Password string
	Username string `gorm:"type:varchar(100);unique_index"`
}

func (user User) String() string {
	return user.FullName
}

func(user *User)BeforeCreate()(err error){
	hashed, err := utils.HashPassword(user.Password)

	user.Password = hashed 

	return
}