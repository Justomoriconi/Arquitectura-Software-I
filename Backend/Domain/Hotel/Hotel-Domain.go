package Domain

type Hotel struct {
	HotelID     int64  `json:"id"`
	Name        string `json:"Name"`
	Rooms       int64  `json:"rooms"`
	Description string `json:"Description"`
}
type Hotels []Hotel
