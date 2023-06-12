package Model

import (
	Hotel "Backend/Model/Hotel"
	User "Backend/Model/User"
)

type Booking struct {
	BookingID int    `gorm:"primary_key;autoIncrement:true;unique"`
	UserID    int    `gorm:"type:integer(255);not null"`
	HotelID   int    `gorm:"type:integer(255);not null"`
	Checkin   string `gorm:"type:Date;not null"`
	Checkout  string `gorm:"type:Date;not null"`

	User  User.User   `gorm:"foreignKey:UserID"`
	Hotel Hotel.Hotel `gorm:"foreignKey:HotelID"`
}
type Bookings []Booking
