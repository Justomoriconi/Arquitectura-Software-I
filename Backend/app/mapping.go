package app

import (
	AvailableordersController "Backend/Controllers/AvailableHotels"
	hotelController "Backend/Controllers/Hotel"
	//	userController "Backend/Controllers/User"
	bookingController "Backend/Controllers/Bookings"
)

// MapUrls maps the urls
func MapUrls() {

	/*	// Users Mapping
		router.GET("/user/:id", userController.GetUserById)
		router.GET("/user", userController.GetUsers)

		// Login Mapping
		router.POST("/login", loginController.Login)
	*/
	// hotel Mapping
	router.GET("/hotel/id/:id", hotelController.GetHotelById)
	router.GET("/hotel/name/:name", hotelController.GetHotelByname)
	// bookings Mapping
	router.GET("/bookings/mybookings/:id", bookingController.GetmyBookings)
	router.GET("/bookings/id/:id", bookingController.GetBookingsById)
	// available Mapping

	router.GET("/availablehotels/:id", AvailableordersController.GetAvailableHotels)
}
