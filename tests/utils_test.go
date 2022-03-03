package main

import (
	"testing"

	"github.com/NonsoAmadi10/echoweb/utils"
	"github.com/stretchr/testify/assert"
)


func TestUtils( t *testing.T){
	t.Run("HashPassword function should has a string of password provided", func(t *testing.T) {
		password:= "hello-123"

		hash, _ := utils.HashPassword(password)

		assert.NotEqual(t, hash, password)
	})

	t.Run("ComparePassword should return true when a hash password compares with its original palintext", func(t *testing.T) {
		password := "yellow"
		fakepwd := "blue"

		hash, _ := utils.HashPassword(password)
		
		compare := utils.CheckPasswordHash(password, hash)

		wrong := utils.CheckPasswordHash(fakepwd, hash)

		assert.Equal(t, true, compare)
		assert.Equal(t, false, wrong)
	})
}