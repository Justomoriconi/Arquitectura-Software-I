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

	log.Debug("User: ", booking)

	return booking, nil
}

func GetBookings() (model.Bookings, error) {
	var bookings model.Bookings
	err := Db.Find(&bookings).Error
	log.Debug("bookings: ", bookings)

	if err != nil {
		log.Println(err)
		return bookings, err
	}

	return bookings, nil
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

func Reserve(userid int, hotelid int, checkin, checkout string) (model.Booking, error) {

	var booking model.Booking

	booking.UserID = userid
	booking.HotelID = hotelid
	booking.Checkin = checkin
	booking.Checkout = checkout

	err := Db.Create(&booking).Error

	if err != nil {
		log.Println(err)
		return booking, err
	}
	return booking, err
}
