package Model

type Hotel struct {
	HotelID int    `gorm:"primaryKey"`
	Name    string `gorm:"type:varchar(255);not null"`
	Rooms   int64  `gorm:"type:integer(255);not null"`
}
type Hotels []Hotel
