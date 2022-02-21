package controllers

import (
	"net/http"
	"time"
	"strconv"
	"github.com/NonsoAmadi10/echoweb/config"
	"github.com/NonsoAmadi10/echoweb/models"
	"github.com/jinzhu/now"
	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
	"github.com/NonsoAmadi10/echoweb/utils"
)

type CreateFlight struct {
	DepartureDate string	`json:"depature_date" validate:"required"`
	ReturnDate	  string	`json:"return_date"`
	DepatureTime string      `json:"departure_time" validate:"required"`
	Origin string 			`json:"origin" validate:"required"`
	Status string 			`json:"status" validate:"required"`
	Destination string 		`json:"destination" validate:"required" `
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

		// format dates
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



		flight.DepartureDate = datatypes.Date(formatdepartureDate)

		if !r.OneWay && r.ReturnDate != "" {
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

			flight.ReturnDate = datatypes.Date(formatreturnDate)
		}

		// time 

		hr, mm, s := utils.TimeFormatter(r.DepatureTime)

		h, _ := strconv.Atoi(hr)
		m, _ := strconv.Atoi(mm)
		ss, _ := strconv.Atoi(s)

		formatTime := datatypes.NewTime(h, m, ss, 0)

		flight.DepatureTime = formatTime
		if err := config.DB.Create(&flight).Error; err != nil {
			return err
		}

		return c.JSONPretty(http.StatusCreated, map[string]models.Flight{"data": flight }, " ")
}