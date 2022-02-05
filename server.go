package main 

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/NonsoAmadi10/echoweb/config"
	"github.com/NonsoAmadi10/echoweb/models"
	"github.com/labstack/echo/v4/middleware"
)

func main(){
	e := echo.New()

	e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"*"},
        AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
    }))

	// Initialize DB
	config.SetupDB(&models.User{})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8081"))
}