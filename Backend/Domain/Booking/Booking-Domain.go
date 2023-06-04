package Domain

type Booking struct {
	BookingID int    `json:"id"`
	UserID    int    `json:"UserID"`
	RoomID    int    `json:"RoomID"`
	HotelID   int    `json:"HotelID"`
	Checkin   string `json:"checkin"`
	Checkout  string `json:"checkout"`
}
