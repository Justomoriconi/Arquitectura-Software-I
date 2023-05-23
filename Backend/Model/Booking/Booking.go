package Model

type Booking struct {
	BookingID uint   `gorm:"primary_key;autoIncrement:true;unique"`
	UserID    int    `gorm:"type:integer(255);not null"`
	HotelID   int    `gorm:"type:integer(255);not null"`
	Day       string `gorm:"type:Date;not null"`
}
type Bookings []Booking
