package main 

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/NonsoAmadi10/echoweb/config"
	"github.com/NonsoAmadi10/echoweb/models"
)

func main(){
	e := echo.New()

	// Initialize DB
	config.SetupDB(&models.User{})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8081"))
}