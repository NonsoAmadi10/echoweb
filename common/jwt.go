package common 
import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/NonsoAmadi10/echoweb/config"
	uuid "github.com/nu7hatch/gouuid"
)

type JwtCustomClaims struct {
	FullName string `json:"name"`
	ID uuid.UUID `json:"id"`
	Email string `json:"email"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func JwtMiddleWare() echo.MiddlewareFunc {
	key := config.GetEnv("JWT_SECRET_KEY")
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(key),
	})
}