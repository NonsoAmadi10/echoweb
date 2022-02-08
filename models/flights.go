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
	DepartureDate time.Time
	ReturnDate	  time.Time
	Origin string `gorm:"type:varchar(100)"`
	Status string `gorm:"type:varchar(100)"`
	Destination string `gorm:"type:varchar(100)"`
	OneWay bool 
	Capacity uint 
	Fare float64
	AirlineCode string `gorm:"type:varchar(100)"`
}

func (f Flight) String() string {
	return f.AirlineCode
}

func (c *Flight)BeforeCreate(tx *gorm.DB) error {
	c.AirlineCode = GenerateSerial(7)
	
	return nil
}

func GenerateSerial(l uint) string {
    buff := make([]byte, int(math.Ceil(float64(l)/float64(1.33333333333))))
	rand.Read(buff)
	str := base64.RawURLEncoding.EncodeToString(buff)
	return str[:l]
}
  