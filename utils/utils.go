package utils

import (
	"os"
	"golang.org/x/crypto/bcrypt"
    _ "github.com/joho/godotenv/autoload"
)


func GetEnv(key string )string{

	  return os.Getenv(key)
}
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}




type Response struct {
	Data interface{} `json:"data"`
	Message string	`json:"message"`
}