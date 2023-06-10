package Clients

import (
	model "Backend/Model/Booking"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetBookingById(id int) (model.Booking, error) { //standard get by id
	var booking model.Booking

	err := Db.Where("booking_id = ?", id).First(&booking).Error

	if err != nil {
		log.Println(err)
		return booking, err
	}

	log.Debug("Booking: ", booking)

	return booking, nil
}

func GetmyBookings(id int) (model.Bookings, error) {
	var booking model.Bookings

	err := Db.Where("user_id = ?", id).Find(&booking).Error
	log.Debug("Bookings: ", booking)

	if err != nil {
		log.Println(err)
		return booking, err
	}

	return booking, err
}

/*
func Reserve(userid ,hotelid , checkin ,checkout string)  error {

	err := Db.Create(b).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return err
} */
