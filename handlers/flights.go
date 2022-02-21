package controllers

import (
	"fmt"
	"net/http"
	_"strconv"
	"time"

	"github.com/NonsoAmadi10/echoweb/config"
	"github.com/NonsoAmadi10/echoweb/models"
	_"github.com/NonsoAmadi10/echoweb/utils"
	"github.com/jinzhu/now"
	"github.com/labstack/echo/v4"
)

type CreateFlight struct {
	DepartureDate string	`json:"departure_date" validate:"required"`
	ReturnDate	  string	`json:"return_date"`
	DepatureTime string      `json:"departure_time" validate:"required"`
	Origin string 			`json:"origin" validate:"required"`
	Status string 			`json:"status"`
	Destination string 		`json:"destination" validate:"required"`
	OneWay bool 			`json:"one_way"` 
	Capacity uint 			`json:"capacity" validate:"required"`
	Fare float64			`json:"fare" validate:"required"`
}

type Response struct {
	Data interface{} `json:"data"`
	Message string	`json:"message"`
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

		//format dates
		t := time.Now()
		 deptDateTime := r.DepartureDate + " " + r.DepatureTime
		formatdepartureDate, err := now.Parse(deptDateTime)
		

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// compare both dates 

		currentTime := formatdepartureDate.After(t)

		if !currentTime {
			return echo.NewHTTPError(http.StatusBadRequest, "departure date cannot be less than current datetime")
		}





		if !r.OneWay {
			if r.ReturnDate == "" {
				return echo.NewHTTPError(http.StatusBadRequest, "Return Date is required")
			}
			returnDateTime := r.ReturnDate + " " + r.ReturnDate

			formatreturnDate, err := now.Parse(returnDateTime)

			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			dateIsNotNow := formatreturnDate.After(t)

			if !dateIsNotNow {
				return echo.NewHTTPError(http.StatusBadRequest, "Return date cannot be less than current datetime")
			}

			dateIsNotDepartDate := formatreturnDate.After(formatdepartureDate)

			if !dateIsNotDepartDate {
				return echo.NewHTTPError(http.StatusBadRequest, "Return time cannot be less than departure datetime")
			}

			flight.ReturnDate = &formatreturnDate
		}


		
		// flight.Status = "scheduled"

		flight = models.Flight{
			DepartureDate: formatdepartureDate,
			DepatureTime: r.DepatureTime,
			OneWay: r.OneWay,
			Capacity: r.Capacity,
			Fare: r.Fare,
			Origin: r.Origin,
			Destination: r.Destination,
			Status: "scheduled",
		}

		if err := config.DB.Create(&flight).Error; err != nil {
			fmt.Println(err.Error())
			return err
		}

		response := &Response{
			Data: &flight,
			Message: "New Flight details added",
		}


		return c.JSONPretty(http.StatusCreated, response, " ")
}