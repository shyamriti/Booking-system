package database

import (
	"Booking-service/pkg/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connection() error {
	var err error
	Db, err = gorm.Open(sqlite.Open("new.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	err = Db.AutoMigrate(&models.Booking{})
	if err != nil {
		return err
	}
	return nil
}
