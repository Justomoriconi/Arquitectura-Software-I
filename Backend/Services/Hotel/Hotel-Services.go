package Services

import (
	client "Backend/Clients/Hotel"
	domain "Backend/Domain/Hotel"
)

type hotelService struct{}

type HotelServiceInterface interface {
	GetHotelById(id int) (domain.Hotel, error)
	GetHotelByname(string) (domain.Hotel, error)
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
	HotelDomain.ID = hotel.HotelID
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
	HotelDomain.ID = hotel.HotelID
	HotelDomain.Rooms = hotel.Rooms

	return HotelDomain, nil
}
