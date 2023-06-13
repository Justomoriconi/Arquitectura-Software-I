package Services

import (
	client "Backend/Clients/AvailableHotels"
	domain "Backend/Domain/Hotel"
)

type availableHotelsService struct{}

type AvailableHotelsServiceInterface interface {
	GetAvailableHotels(string, string) (domain.Hotels, error)
}

var (
	HotelServices AvailableHotelsServiceInterface
)

func init() {
	HotelServices = &availableHotelsService{}
}

/*
func (s *availableHotelsService) GetAvailableHotels(checkin, checkout string) (domain.Hotels, error) {

	hotel, err := client.GetAvailableHotels(checkin, checkout)
	var hotelDomains domain.Hotels

	if err != nil {
		return hotelDomains, err
	}
	for _, hotel1 := range hotel {
		var hotelDomain domain.Hotel
		hotelDomain.Name = hotel1.Name
		hotelDomain.HotelID = hotel1.HotelID
		hotelDomain.Rooms = hotel1.Rooms
		hotelDomain.Description = hotel1.Description
		hotelDomains = append(hotelDomains, hotelDomain)
	}

	return hotelDomains, nil
}*/

func (s *availableHotelsService) GetAvailableHotels(Checkin, Checkout string) (domain.Hotels, error) {
	hotels, err := client.GetAvailableHotels(Checkin, Checkout)
	if err != nil {
		return nil, err
	}
	var hotelDomains domain.Hotels
	for _, hotel := range hotels {
		hotelDomain := domain.Hotel{
			Name:        hotel.Name,
			HotelID:     hotel.HotelID,
			Rooms:       hotel.Rooms,
			Description: hotel.Description,
		}
		hotelDomains = append(hotelDomains, hotelDomain)
	}
	return hotelDomains, nil
}
