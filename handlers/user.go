package controllers

import (
	"fmt"
	"net/http"

	"github.com/NonsoAmadi10/echoweb/config"
	"github.com/NonsoAmadi10/echoweb/models"
	"github.com/NonsoAmadi10/echoweb/utils"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
)

type NewUser struct {
	Fullname string `json:"fullname" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Username string `json:"username" validate:"required"`
}

type LogUser struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5"`
}

type Data struct {
	ID uuid.UUID `json:"id"`
	Email string `json:"email"`
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

func LoginUser(c echo.Context)(err error){
	request := new(LogUser)

	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(request); err != nil {
		return err
	}


	var user models.User
	existingUser := config.DB.First(&user, "email = ?", request.Email)

	if existingUser.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid email or password")
	}

	if matched:= utils.CheckPasswordHash(request.Password ,user.Password); !matched {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid email or password")
	}
	
	response:= &Data{ID: user.ID, Email: user.Email}
	
	return c.JSONPretty(http.StatusOK, response, " ")


}