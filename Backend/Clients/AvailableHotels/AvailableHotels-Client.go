package Clients

import (
	model "Backend/Model/Hotel"

	Db "Backend/Database"
)

func GetAvailableHotels(checkIn string, checkOut string) (model.Hotels, error) {
	var hotels model.Hotels
	db := Db.DbEngine()
	err := db.Select("DISTINCT(hotels.hotel_id), hotels.name,hotels.rooms, hotels.description").
		Table("hotels").
		Where(`
			hotels.rooms > (
				SELECT COUNT(*)
				FROM bookings
				WHERE bookings.hotel_id = hotels.hotel_id AND (
					(bookings.checkin >= ? AND bookings.checkin <= ?) OR
					(bookings.checkout >= ? AND bookings.checkout <= ?) OR
					(bookings.checkin <= ? AND bookings.checkout >= ?)
				)
			)
		`, checkIn, checkOut, checkIn, checkOut, checkIn, checkOut).
		Scan(&hotels).Error
	if err != nil {
		return nil, err
	}

	return hotels, nil
}
