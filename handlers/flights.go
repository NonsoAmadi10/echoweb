package controllers

import (
	"net/http"
	"github.com/NonsoAmadi10/echoweb/config"
	"github.com/NonsoAmadi10/echoweb/models"
	"github.com/labstack/echo/v4"
)

type CreateFlight struct {
	DepartureDate string	`json:"depature_date" validate:"required"`
	ReturnDate	  string	`json:"return_date" validate:"required"`
	DepatureTime string      `json:"departure_time"`
	Origin string 			`json:"origin" validate:"required"`
	Status string 			`json:"status" validate:"required"`
	Destination string 		`json:"destination" validate:"required"`
	OneWay bool 			`json:"oneway"`
	Capacity uint 			`json:"capacity" validate:"required"`
	Fare float64			`json:"fare" validate:"required"`
}


func AddFlight(c echo.Context) (err error) {

		r := new(CreateFlight)

		if err := c.Bind(&r); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err = c.Validate(r); err != nil {
			return err
		}

		var flight models.Flight

		if err := config.DB.Create(&flight).Error; err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, "done")
}