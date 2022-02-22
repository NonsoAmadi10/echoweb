package controllers

import (
	"fmt"
	"net/http"
	"github.com/NonsoAmadi10/echoweb/config"
	"github.com/NonsoAmadi10/echoweb/models"
	"github.com/NonsoAmadi10/echoweb/utils"
	 "github.com/NonsoAmadi10/echoweb/helpers"
	"github.com/labstack/echo/v4"
)






func AddFlight(c echo.Context) (err error) {

		r := new(helpers.CreateFlight)


		if err := c.Bind(&r); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err = c.Validate(r); err != nil {
			return err
		}

		var flight models.Flight

		//format dates
		formatdepartureDate, returnDate, err := helpers.TimeFormatter(r)
		
		// catch err 

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		flight = models.Flight{
			DepartureDate: formatdepartureDate,
			DepatureTime: r.DepatureTime,
			OneWay: r.OneWay,
			Capacity: r.Capacity,
			Fare: r.Fare,
			Origin: r.Origin,
			Destination: r.Destination,
			Status: "scheduled",
			ReturnDate: returnDate,
		}

		if err := config.DB.Create(&flight).Error; err != nil {
			fmt.Println(err.Error())
			return err
		}

		response := &utils.Response{
			Data: &flight,
			Message: "New Flight details added",
		}


		return c.JSONPretty(http.StatusCreated, response, " ")
}