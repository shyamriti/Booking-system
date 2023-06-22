package router

import (
	"Booking-service/pkg/controller"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()

	r.GET("/seats", controller.GetAllSeats)
	r.GET("/seats/:id", controller.GetSeatPricing)
	r.POST("/booking", controller.CreateBooking)
	r.GET("/booking/:useridentifier", controller.RetrieveBooking)
	return r
}
