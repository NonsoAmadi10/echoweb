package controllers

import (
	"fmt"
	"net/http"

	"github.com/NonsoAmadi10/echoweb/config"
	"github.com/NonsoAmadi10/echoweb/helpers"
	"github.com/NonsoAmadi10/echoweb/models"
	"github.com/NonsoAmadi10/echoweb/utils"
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
		DepatureTime:  r.DepatureTime,
		OneWay:        r.OneWay,
		Capacity:      r.Capacity,
		Fare:          r.Fare,
		Origin:        r.Origin,
		Destination:   r.Destination,
		Status:        "scheduled",
		ReturnDate:    returnDate,
	}

	if err := config.DB.Create(&flight).Error; err != nil {
		fmt.Println(err.Error())
		return err
	}

	response := &utils.Response{
		Data:    &flight,
		Message: "New Flight details added",
	}

	return c.JSONPretty(http.StatusCreated, response, " ")
}

func GetAllFlights(c echo.Context) (err error) {
	var flights []models.Flight

	config.DB.Find(&flights)

	response := &utils.Response{
		Data:    &flights,
		Message: "All Flights Details",
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}

func GetAvailableFlights(c echo.Context) (err error) {
	var flights []models.Flight

	config.DB.Where("status = ?", "scheduled").Find(&flights)

	response := &utils.Response{
		Data:    &flights,
		Message: "All Flights Details",
	}
	return c.JSONPretty(http.StatusOK, response, " ")
}

func GetFlightInfo(c echo.Context) (err error) {
	id := c.Param("id")

	var flight models.Flight

	config.DB.First(&flight, "id = ?", id)

	response := &utils.Response{
		Data:    &flight,
		Message: "flight details returned successfully",
	}
	return c.JSONPretty(http.StatusOK, response, " ")
}

func UpdateFlightInfo(c echo.Context) (err error) {

	id := c.Param("id")
	r := new(helpers.UpdateFlight)

	var flight models.Flight
	if err := c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if r.Status != "cancelled" || r.Status != "completed" {
		return echo.NewHTTPError(http.StatusBadRequest, "status can only be updated to completed or cancelled")
	}
	if err := config.DB.Where("id = ?", id).First(&flight).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "record not found!")
	}

	config.DB.Model(&flight).Updates(models.Flight{Status: r.Status})

	response := &utils.Response{
		Data:    &flight,
		Message: "flight details updated successfully",
	}
	return c.JSONPretty(http.StatusOK, response, " ")

}
