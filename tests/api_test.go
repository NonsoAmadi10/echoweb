package main

import (
	// "os"
	"testing"
	_ "github.com/joho/godotenv/autoload"
	"github.com/NonsoAmadi10/echoweb/app"
	"github.com/steinfletcher/apitest"
	"github.com/steinfletcher/apitest-jsonpath"
)

func TestMain(t *testing.T) {
	// os.Setenv("DATABASE_URL", "postgres://ikayaocf:fvQNB0k_X3eNS2ZIPouyENKaFOCVU4pU@tyke.db.elephantsql.com/ikayaocf")

	t.Run("It should register a new User", func(t *testing.T) {
		apitest.New().
			Handler(app.StartApp()).
			Post("/api/v1/register").
			JSON(`{ "fullname": "Amadi Chinonso", "email": "nonsoamadi@aol.com", "password": "chivulena", "role": "customer","username": "holo12"}`).
			Expect(t).
			Assert(jsonpath.Present("message")).
			Assert(jsonpath.Equal("$.message", "nonsoamadi@aol.com has been successfully created")).
			End()
	})

}
