package models

import (
	"fmt"
	"time"
	Common "github.com/NonsoAmadi10/echoweb/common"
	"github.com/NonsoAmadi10/echoweb/config"
	"github.com/NonsoAmadi10/echoweb/utils"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type User struct {
    Common.Model
	Email string `gorm:"type:varchar(100);unique_index"`
	FullName     string
	Password string
	Username string `gorm:"type:varchar(100);unique_index"`
}

func (user User) String() string {
	return user.FullName
}

func(user *User)BeforeCreate(tx *gorm.DB)(err error){
	hashed, err := utils.HashPassword(user.Password)
	user.Password = hashed 

	return
}

func GenerateJWT(user *User) (string, error) {
	var mySigningKey = config.GetEnv("JWT_SECRET_KEY")
	claims := &Common.JwtCustomClaims{
		ID: user.ID,
        FullName: user.FullName,
        Email: user.Email,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
        },
    }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(mySigningKey)



	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return t, nil
}
