package models

import (
	"time"

	Common "github.com/NonsoAmadi10/echoweb/common"
	"github.com/NonsoAmadi10/echoweb/config"
	"github.com/NonsoAmadi10/echoweb/utils"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type User struct {
    Common.Model
	Email string `gorm:"type:varchar(100);uniqueIndex"`
	FullName     string
	Password string
	Username string `gorm:"type:varchar(100);unique;not_null"`
	Role string 
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
	var key = config.GetEnv("JWT_SECRET_KEY")
	claims := &Common.JwtCustomClaims{
		ID: user.ID,
        FullName: user.FullName,
        Email: user.Email,
		Username: user.Username,
		Role: user.Role,
		StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
        },
    }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(key))



	if err != nil {
		log.Info(err)
		return "", err
	}
	return t, nil
}
