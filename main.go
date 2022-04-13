package main

import (
	"log"

	App "github.com/NonsoAmadi10/echoweb/app"
)

func main() {

	err := App.StartApp().Start("localhost:8082")
	if err != nil {
		log.Fatal(err)
	}
}
