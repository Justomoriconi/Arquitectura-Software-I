package Services

import (
	client "Backend/Clients/Hotel"
	domain "Backend/Domain/Hotel"
)

type hotelService struct{}

type HotelServiceInterface interface {
	GetHotelById(id int) (domain.Hotel, error)
	GetHotelByname(string) (domain.Hotel, error)
	GetHotels() (domain.Hotels, error)
}

var (
	HotelService HotelServiceInterface
)

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) GetHotelById(id int) (domain.Hotel, error) {

	hotel, err := client.GetHotelById(id)
	var HotelDomain domain.Hotel

	if err != nil {
		return HotelDomain, err
	}
	if hotel.HotelID == 0 {
		return HotelDomain, nil
	}
	HotelDomain.Name = hotel.Name
	HotelDomain.HotelID = hotel.HotelID
	HotelDomain.Rooms = hotel.Rooms

	return HotelDomain, nil
}

func (s *hotelService) GetHotelByname(name string) (domain.Hotel, error) {

	hotel, err := client.GetHotelByName(name)
	var HotelDomain domain.Hotel

	if err != nil {
		return HotelDomain, err
	}
	if hotel.HotelID == 0 {
		return HotelDomain, nil
	}
	HotelDomain.Name = hotel.Name
	HotelDomain.HotelID = hotel.HotelID
	HotelDomain.Rooms = hotel.Rooms
	HotelDomain.Description = hotel.Description
	return HotelDomain, nil
}
func (s *hotelService) GetHotels() (domain.Hotels, error) {

	hotel, err := client.GetHotels()
	var hotelDomains domain.Hotels

	if err != nil {
		return hotelDomains, err
	}
	for _, hotel := range hotel {
		var hotelDomain domain.Hotel
		hotelDomain.Name = hotel.Name
		hotelDomain.HotelID = hotel.HotelID
		hotelDomain.Rooms = hotel.Rooms
		hotelDomain.Description = hotel.Description
		hotelDomains = append(hotelDomains, hotelDomain)
	}

	return hotelDomains, nil
}
