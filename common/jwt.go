package common

import (

	"github.com/NonsoAmadi10/echoweb/config"
	"github.com/golang-jwt/jwt"
	uuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	key := config.GetEnv("JWT_SECRET_KEY")
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