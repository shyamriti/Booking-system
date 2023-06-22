package controller

import (
	"Booking-service/pkg/database"
	"Booking-service/pkg/models"
	"strconv"
	"strings"
)

func calculatePrice(count int64, v uint) (float64, error) {
	var (
		seats       models.Seats
		seatPricing models.SeatPricing
		err         error
	)

	err = database.Db.Where("seat_id= ?", v).First(&seats).Error
	if err != nil {
		return 0, err
	}

	err = database.Db.Raw("select * from seat_pricings where seat_class=?", seats.SeatClass).First(&seatPricing).Error
	if err != nil {
		return 0, err
	}

	seatPricingSting := findSeatPrice(count, seatPricing)
	seatPrice, err := convertPriceStrFloat(seatPricingSting)
	if err != nil {
		return 0, err
	}
	return seatPrice, nil

}

func findSeatPrice(count int64, seatPricing models.SeatPricing) string {
	var price string

	bookedSeatPercentage := (count / 500) * 100

	switch {
	case bookedSeatPercentage >= 40 && bookedSeatPercentage <= 60:
		if seatPricing.NormalPrice == "" {
			price = seatPricing.MaxPrice
		} else {
			price = seatPricing.NormalPrice
		}

	case bookedSeatPercentage < 40:
		if seatPricing.MinPrice == "" {
			price = seatPricing.NormalPrice
		} else {
			price = seatPricing.MinPrice
		}

	case bookedSeatPercentage > 60:
		if seatPricing.MaxPrice == "" {
			price = seatPricing.NormalPrice
		} else {
			price = seatPricing.MaxPrice
		}

	}
	return price
}

func strRemoveAt(s string) string {
	s = strings.ReplaceAll(s, "$", "")
	return s
}

func convertPriceStrFloat(priceAsString string) (float64, error) {
	seatPriceAsString := strRemoveAt(priceAsString)
	seatPrice, err := strconv.ParseFloat(seatPriceAsString, 64)
	if err != nil {
		return 0, err
	}
	return seatPrice, nil
}
