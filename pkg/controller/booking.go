package controller

import (
	"Booking-service/pkg/database"
	"Booking-service/pkg/models"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateBooking(c *gin.Context) {
	var (
		bookings    models.Bookings
		booking     models.Booking
		seatId      uint
		totalAmount float64
		count       int64
		err         error
	)

	err = database.Db.Model(&booking).Where("is_booked = ?", true).Count(&count).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to count booked seats"})
		return
	}

	err = c.ShouldBindJSON(&bookings)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to binding JSON"})
		return
	}
	bookingId := uuid.NewString()

	result := database.Db.Raw("select bookings.seat_id from bookings left join seats on seats.seat_id = bookings.seat_id where bookings.seat_id in (?)", bookings.SeatIds).First(&seatId)
	if result.Error == gorm.ErrRecordNotFound {
		for _, v := range bookings.SeatIds {
			booking = models.Booking{
				BookingId:   bookingId,
				Name:        bookings.Name,
				Email:       bookings.Email,
				PhoneNumber: bookings.PhoneNumber,
				SeatId:      v,
				IsBooked:    true,
			}

			price, err := calculatePrice(count, v)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to calculate price of seat"})
			}

			err = database.Db.Create(&booking).Error
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to create booking"})
				return
			}
			totalAmount += price
		}

	} else if seatId > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "seat not available")
		// msg := "This seats are previously booked, Please select another."
		// SendBookingNotification(bookings.PhoneNumber, msg)
		return
	}
	// msg := "your seat is booked"
	// SendBookingNotification(bookings.PhoneNumber, msg)
	c.JSON(200, gin.H{
		"booking_id":   bookingId,
		"total_amount": totalAmount,
	})
}

func RetrieveBooking(c *gin.Context) {
	var booking []models.Booking
	var value string
	userIdentifier := c.Param("useridentifier")
	if userIdentifier == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Please provide valid user identifier")
		return
	} else {
		value = CheckEmailOrPhoneNumber(userIdentifier)
		if value == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, "Please enter valid user identifier")
			return
		}
	}

	err := database.Db.Where(value, userIdentifier).Find(&booking).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Error")
		return
	}
	c.JSON(200, booking)
}

func CheckEmailOrPhoneNumber(userIdentifier string) string {
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	phonePattern := `^\+?[1-9]\d{1,14}$`

	emailReg := regexp.MustCompile(emailPattern)
	phoneReg := regexp.MustCompile(phonePattern)

	if emailReg.MatchString(userIdentifier) {
		return "email=?"
	}

	if phoneReg.MatchString(userIdentifier) {
		return "phone_number=?"
	}

	return ""
}
