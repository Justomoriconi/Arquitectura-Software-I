package Services

import (
	client "Backend/Clients/Booking"
	domain "Backend/Domain/Booking"
)

type bookingService struct{}

type BookingServiceInterface interface {
	GetBookinglById(id int) (domain.Booking, error)
	GetmyBookings(id int) (domain.Bookings, error)
}

var (
	BookingService BookingServiceInterface
)

func init() {
	BookingService = &bookingService{}
}

func (s *bookingService) GetBookinglById(id int) (domain.Booking, error) {

	booking, err := client.GetBookingById(id)
	var BookingDomain domain.Booking

	if err != nil {
		return BookingDomain, err
	}
	if booking.HotelID == 0 {
		return BookingDomain, nil
	}
	BookingDomain.BookingID = booking.BookingID
	BookingDomain.HotelID = booking.HotelID
	BookingDomain.UserID = booking.UserID
	BookingDomain.Checkin = booking.Checkin
	BookingDomain.Checkout = booking.Checkout

	return BookingDomain, nil
}

func (s *bookingService) GetmyBookings(id int) (domain.Bookings, error) {

	booking, err := client.GetmyBookings(id)
	var bookingsDomains domain.Bookings

	if err != nil {
		return bookingsDomains, err
	}

	for _, booking := range booking {
		var BookingDomain domain.Booking
		BookingDomain.BookingID = booking.BookingID
		BookingDomain.HotelID = booking.HotelID
		BookingDomain.UserID = booking.UserID
		BookingDomain.Checkin = booking.Checkin
		BookingDomain.Checkout = booking.Checkout

		bookingsDomains = append(bookingsDomains, BookingDomain)
	}

	return bookingsDomains, nil
}
