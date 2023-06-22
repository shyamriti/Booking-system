package models

import "gorm.io/gorm"

type SeatPricing struct {
	Id          uint   `json:"id"`
	SeatClass   string `json:"seat_class"`
	MinPrice    string `json:"min_price"`
	NormalPrice string `json:"normal_price"`
	MaxPrice    string `json:"max_price"`
}

type Seats struct {
	SeatId         uint   `json:"seat_id"`
	SeatIdentifier string `json:"seat_identifier"`
	SeatClass      string `json:"seat_class"`
}

type Booking struct {
	gorm.Model
	BookingId   string `json:"booking_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	SeatId      uint   `json:"seat_id"`
	IsBooked    bool   `json:"is_booked"`
}

type Bookings struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	SeatIds     []uint `json:"seat_id"`
}
