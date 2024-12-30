package main

import (
	apps "auth-server/internal/app"
	"log"
)

func main() {

	// Entry point of the application
	apps.StartServer()
	log.Println("Application up and running")

}
