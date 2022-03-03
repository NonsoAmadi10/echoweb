package common

import (
	"time"

	"github.com/NonsoAmadi10/echoweb/models"
	"github.com/NonsoAmadi10/echoweb/utils"
	"github.com/golang-jwt/jwt"
	uuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type JwtCustomClaims struct {
	FullName string `json:"name"`
	ID uuid.UUID `json:"id"`
	Email string `json:"email"`
	Username string `json:"username"`
	Role string `json:"role"`
	jwt.StandardClaims
}

func JwtMiddleWare() echo.MiddlewareFunc {
	key := utils.GetEnv("JWT_SECRET_KEY")
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(key),
	})
}

func ServerAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JwtCustomClaims)
		if claims.Role != "admin" {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

func GenerateJWT(user *models.User) (string, error) {
	var key = utils.GetEnv("JWT_SECRET_KEY")
	claims := &JwtCustomClaims{
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