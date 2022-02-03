package config

import (
	"log"
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB 

func GetEnv(key string )string{

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	  }
	
	  return os.Getenv(key)
}


func SetupDB(model...interface{}){

	dbString := GetEnv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{})


	if err != nil {
		panic("Failed to Connect to Database")
	}

	db.AutoMigrate(model...)

	DB = db
}
