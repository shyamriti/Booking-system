package controller

import (
	"Booking-service/pkg/database"
	"Booking-service/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllSeats(c *gin.Context) {
	var (
		seats    []models.Seats
		seatId   uint
		isBooked bool
	)

	if err := database.Db.Order("seat_class asc").Find(&seats).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve seats"})
		return
	}

	for i, v := range seats {
		seatId = v.SeatId
		err := database.Db.Raw("select is_booked from bookings where seat_id=?", seatId).First(&isBooked).Error
		if err != nil {
			isBooked = false
		}
		c.JSON(200, gin.H{
			"seats":     seats[i],
			"is_booked": isBooked,
		})
	}
}

func GetSeatPricing(c *gin.Context) {
	var (
		seats       models.Seats
		seatPricing models.SeatPricing
		count       int64
		booking     models.Booking
		err         error
		seatId      int
	)

	seatIdAsString := c.Param("id")
	seatId, err = strconv.Atoi(seatIdAsString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "string is nil")
		return
	}

	err = database.Db.Where("seat_id= ?", seatId).First(&seats).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve seats"})
		return
	}

	err = database.Db.Where("seat_class=?", seats.SeatClass).First(&seatPricing).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve seat pricing"})
		return
	}

	err = database.Db.Model(&booking).Where("is_booked = ?", true).Count(&count).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to count booked seats"})
		return
	}

	price := findSeatPrice(count, seatPricing)
	c.JSON(200, gin.H{
		"seats": seats,
		"price": price,
	})
}
