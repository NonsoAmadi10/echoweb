package controllers

import (
	"net/http"
	"github.com/NonsoAmadi10/echoweb/config"
	"github.com/NonsoAmadi10/echoweb/models"
	"github.com/NonsoAmadi10/echoweb/utils"
	"github.com/labstack/echo/v4"
	"github.com/NonsoAmadi10/echoweb/common"
)

type NewUser struct {
	Fullname string `json:"fullname" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Username string `json:"username" validate:"required"`
	Role string `json:"role" default:"customer"`
}

type LogUser struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5"`
}

type Data struct {
	Token string `json:"token"`
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
		return echo.NewHTTPError(http.StatusBadRequest, "email is already used")
	  }
	  user = models.User{
		Email: reqBody.Email,
		FullName: reqBody.Fullname,
		Username: reqBody.Username,
		Password: reqBody.Password,
		Role: reqBody.Role,
	  }
	

	if err := config.DB.Create(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}
	msg := user.Email + " " + "has been successfully created"
	response := &utils.Response{
		Message: msg,
		Data: map[string]string{},
	}
	return c.JSONPretty(http.StatusCreated, response, " ")
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
	
	t, err := common.GenerateJWT(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response:= &utils.Response{
		Data: &Data{ Token: t },
		Message: "login successful!",
	}
	
	return c.JSONPretty(http.StatusOK, response, " ")


}