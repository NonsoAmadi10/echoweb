package main

import (
	"net/http"

	"github.com/NonsoAmadi10/echoweb/config"
	"github.com/NonsoAmadi10/echoweb/handlers"
	"github.com/NonsoAmadi10/echoweb/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
    validator *validator.Validate
  }

  func (cv *CustomValidator) Validate(i interface{}) error {
    if err := cv.validator.Struct(i); err != nil {
      // Optionally, you could return the error to give each route more control over the status code
      return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return nil
  }

func main(){
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"*"},
        AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
    }))

	// Initialize DB
	config.SetupDB(&models.User{}, &models.Flight{})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	api := e.Group("/api/v1")
	api.GET("/register", controllers.RegisterUser)
	api.POST("/login", controllers.LoginUser)

	// Admin Routers
	admin := api.Group("/flights")
	admin.POST("", controllers.AddFlight)
	
	e.Logger.Fatal(e.Start(":8081"))
}