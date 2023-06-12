package Clients

import (
	model "Backend/Model/Booking"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

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

	var bookings model.Booking

	err := Db.
		Table("hotels").
		Select("hotels.hotel_id, hotels.name").
		Joins("LEFT JOIN bookings ON bookings.hotel_id = hotels.hotel_id").
		Where("hotels.hotel_id = ? AND hotel.rooms - 1 > (SELECT COUNT(*) FROM bookings WHERE bookings.hotel_id = hotels.hotel_id AND ? < bookings.checkout AND ? > bookings.checkin OR bookings.hotel_id IS NULL)", hotelid, checkout, checkin).
		Scan(&bookings).Error

	if err != nil {
		return bookings, err
	}
	var nbookings model.Booking
	nbookings.UserID = userid
	nbookings.HotelID = hotelid
	nbookings.Checkin = checkin
	nbookings.Checkout = checkout

	error := Db.Create(nbookings).Error
	if error != nil {
		log.Println(err)
		return bookings, error
	}
	return nbookings, error
}
