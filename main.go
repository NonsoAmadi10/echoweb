package main

import (
	"log"

	App "github.com/NonsoAmadi10/echoweb/app"
	
)



func main(){
	
	err := App.StartApp().Start("localhost:8081")
	if err != nil {
		log.Fatal(err)
	}
}