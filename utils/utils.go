package utils

import (
	"strings"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func TimeFormatter(tyme string)(h string, m string, s string) {
    t := strings.Split(tyme, ":")
    return t[0], t[1], t[2]
}

