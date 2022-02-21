package models

import (
	"crypto/rand"
	"encoding/base64"
	"math"
	"time"

	Common "github.com/NonsoAmadi10/echoweb/common"
	"gorm.io/gorm"
)



type Flight struct {
	Common.Model
	DepartureDate time.Time `json:"departure_date"`
	ReturnDate	  *time.Time `json:"return_date" gorm:"type:TIMESTAMP NULL"` 
	DepatureTime string `json:"departure_time"`
	Origin string `json:"origin" gorm:"type:varchar(100)"`
	Status string `json:"status" gorm:"type:varchar(100)"`
	Destination string `json:"destination" gorm:"type:varchar(100)"`
	OneWay bool `json:"one_way"`
	Capacity uint `json:"capacity"`
	Fare float64	`json:"fare"`
	AirlineCode string `json:"airline_code" gorm:"type:varchar(100)"`
}

func (f Flight) String() string {
	return f.AirlineCode
}

func (c *Flight)BeforeSave(tx *gorm.DB) error {
	c.AirlineCode = GenerateSerial(7)
	
	return nil
}

func GenerateSerial(l uint) string {
    buff := make([]byte, int(math.Ceil(float64(l)/float64(1.33333333333))))
	rand.Read(buff)
	str := base64.RawURLEncoding.EncodeToString(buff)
	return str[:l]
}
  