package controllers

import (
	"fmt"
	"net/http"

	"github.com/NonsoAmadi10/echoweb/config"
	"github.com/NonsoAmadi10/echoweb/models"
	"github.com/labstack/echo/v4"
)

type NewUser struct {
	Fullname string `json:"fullname" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Username string `json:"username" validate:"required"`
}

func RegisterUser(c echo.Context)(err error){ 
	reqBody := new(NewUser)
	if err := c.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(reqBody); err != nil {
		return err
	}

	var user models.User
	// find user 
	
	if err := config.DB.Find(&user, "email = ?", reqBody.Email).RowsAffected; err > 0 {
		// error handling...
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "email is already used")
	  }
	  user = models.User{
		Email: reqBody.Email,
		FullName: reqBody.Fullname,
		Username: reqBody.Username,
		Password: reqBody.Password,
	  }
	

	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}
	response := user.Email + " " + "has been successfully created"
	return c.JSON(http.StatusCreated, response)
}