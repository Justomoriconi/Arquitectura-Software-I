package Services

import (
	client "Backend/Clients/Booking"
	domain "Backend/Domain/Booking"
)

type bookingService struct{}

type BookingServiceInterface interface {
	GetmyBookings(id int) (domain.Bookings, error)
	GetBookings() (domain.Bookings, error)
	GetBookingsbyid(id int) (domain.Booking, error)
	Reserve(id, hotelid int, string2 string, string3 string) (domain.Booking, error)
}

var (
	BookingService BookingServiceInterface
)

func init() {
	BookingService = &bookingService{}
}
func (s *bookingService) GetBookings() (domain.Bookings, error) {
	booking, err := client.GetBookings()
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
func (s *bookingService) GetBookingsbyid(id int) (domain.Booking, error) {
	booking, err := client.GetBookingById(id)
	var bookingsDomains domain.Booking

	if err != nil {
		return bookingsDomains, err
	}

	var BookingDomain domain.Booking

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

func (s *bookingService) Reserve(id int, hotelid int, checkin string, checkout string) (domain.Booking, error) {
	var BookingDomain domain.Booking

	booking, err := client.Reserve(id, hotelid, checkin, checkout)

	if err != nil {
		return BookingDomain, err
	}

	BookingDomain.BookingID = booking.BookingID
	BookingDomain.HotelID = booking.HotelID
	BookingDomain.UserID = booking.UserID
	BookingDomain.Checkin = booking.Checkin
	BookingDomain.Checkout = booking.Checkout

	return BookingDomain, nil
}
