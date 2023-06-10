package Model

type Hotel struct {
	HotelID     int64  `gorm:"primary_key;autoIncrement:true;unique"`
	Name        string `gorm:"type:varchar(255);not null"`
	Rooms       int64  `gorm:"type:integer(255);not null"`
	Description string `gorm:"type:varchar(255);not null"`
}
type Hotels []Hotel
