package helpers

import (
	"net/http"
	"time"

	"github.com/jinzhu/now"
	"github.com/labstack/echo"
)

type CreateFlight struct {
	DepartureDate string  `json:"departure_date" validate:"required"`
	ReturnDate    string  `json:"return_date"`
	DepatureTime  string  `json:"departure_time" validate:"required"`
	Origin        string  `json:"origin" validate:"required"`
	Status        string  `json:"status"`
	Destination   string  `json:"destination" validate:"required"`
	OneWay        bool    `json:"one_way"`
	Capacity      uint    `json:"capacity" validate:"required"`
	Fare          float64 `json:"fare" validate:"required"`
}

type UpdateFlight struct {
	Status string `json:"status" validate:"required"`
}

func TimeFormatter(r *CreateFlight) (formatdepartureDate time.Time, returnDate *time.Time, err error) {
	time.Now()
	t := time.Now()
	deptDateTime := r.DepartureDate + " " + r.DepatureTime
	formatdepartureDate, err = now.Parse(deptDateTime)

	if err != nil {
		return t, &t, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// compare both dates

	currentTime := formatdepartureDate.After(t)

	if !currentTime {
		return t, &t, echo.NewHTTPError(http.StatusBadRequest, "departure date cannot be less than current datetime")
	}

	if !r.OneWay {
		if r.ReturnDate == "" {
			return t, &t, echo.NewHTTPError(http.StatusBadRequest, "Return Date is required")
		}
		returnDateTime := r.ReturnDate + " " + r.DepatureTime

		formatreturnDate, err := now.Parse(returnDateTime)

		if err != nil {
			return t, &t, echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		dateIsNotNow := formatreturnDate.After(t)

		if !dateIsNotNow {
			return t, &t, echo.NewHTTPError(http.StatusBadRequest, "Return date cannot be less than current datetime")
		}

		dateIsNotDepartDate := formatreturnDate.After(formatdepartureDate)

		if !dateIsNotDepartDate {
			return t, &t, echo.NewHTTPError(http.StatusBadRequest, "Return time cannot be less than departure datetime")
		}
		returnDate = &formatreturnDate
	}

	return formatdepartureDate, returnDate, nil
}
