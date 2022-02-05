package common

import (
	// "gorm.io/gorm"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Base Struct will contain columns that cut across all tables
type Model struct{
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" sql:"default:uuid.NewV4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}


// Set Primary Key ID as UUID while saving to database
func (b *Model)BeforeSave(tx *gorm.DB)( error){
	u, err := uuid.NewV4()
	if err != nil {
		return err
	}

	tx.Statement.SetColumn("ID", u)
	return nil
}