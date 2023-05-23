package Domain

type Booking struct {
	BookingID int    `json:"id"`
	UserID    int    `json:"UserID"`
	HotelID   int    `json:"HotelID"`
	Date      string `json:"Date"`
}
