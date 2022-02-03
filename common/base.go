package common 

import (
	"gorm.io/gorm"
	"github.com/nu7hatch/gouuid"
	"time"
)

// Base Struct will contain columns that cut across all tables 
type Base struct{
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

// Set Primary Key ID as UUID while saving to database
func (base *Base)BeforeCreate(scope *gorm.DB)(error){
	u, err := uuid.NewV4()
	if err != nil {
		return err
	}

	scope.Statement.SetColumn("ID", u)

	return nil
}