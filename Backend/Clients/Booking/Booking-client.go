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

	err := Db.Where("booking_id = ?", id).Find(&booking).Error
	log.Debug("Bookings: ", booking)

	if err != nil {
		log.Println(err)
		return booking, err
	}

	return booking, err
}

func Reserve(userid int, hotelid int, checkin, checkout string) (model.Booking, error) {

	var booking model.Booking

	err := Db.
		Table("hotels").
		Select("hotels.hotel_id, hotels.name").
		Joins("LEFT JOIN bookings ON bookings.hotel_id = hotels.hotel_id").
		Where("hotels.hotel_id = ? AND hotel.rooms - 1 > (SELECT COUNT(*) FROM bookings WHERE bookings.hotel_id = hotels.hotel_id AND ? < bookings.checkout AND ? > bookings.checkin OR bookings.hotel_id IS NULL)", hotelid, checkout, checkin).
		Scan(&booking).Error

	if err != nil {
		return booking, err
	}

	booking.UserID = userid
	booking.HotelID = hotelid
	booking.Checkin = checkin
	booking.Checkout = checkout

	error := Db.Create(booking).Error
	if error != nil {
		log.Println(err)
		return booking, error
	}
	return booking, error
}
