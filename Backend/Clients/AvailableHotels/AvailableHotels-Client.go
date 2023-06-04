package Clients

import (
	model "Backend/Model/Hotel"

	Db "Backend/Database"
)

func GetAvailableHotels(checkIn string, checkOut string) (model.Hotels, error) {
	var hotels model.Hotels
	db := Db.DbEngine()
	err := db.
		Select("DISTINCT(hotels.hotel_id), hotels.name").
		Table("hotels").
		Joins("LEFT JOIN bookings ON bookings.hotel_id = hotels.hotel_id").
		Where("hotels.rooms > (SELECT COUNT(*) FROM bookings WHERE bookings.hotel_id = hotels.hotel_id AND ? < bookings.checkout AND ? > bookings.checkin OR bookings.hotel_id IS NULL)", checkOut, checkIn).
		Scan(&hotels).Error
	if err != nil {
		return nil, err
	}

	return hotels, nil
}
