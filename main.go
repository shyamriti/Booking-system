package main

import (
	"Booking-service/pkg/database"
	"Booking-service/pkg/router"
	"log"
)

func main() {
	err := database.Connection()
	if err != nil {
		log.Fatal(err)
	}
	r := router.Routes()

	r.Run(":8080")
}
