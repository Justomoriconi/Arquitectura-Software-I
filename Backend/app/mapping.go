package app

import (
	auth "Backend/Controllers/Auth"
	AvailableordersController "Backend/Controllers/AvailableHotels"
	bookingController "Backend/Controllers/Bookings"
	hotelController "Backend/Controllers/Hotel"
	userController "Backend/Controllers/User"
)

// MapUrls maps the urls
func MapUrls() {

	/*	// Users Mapping
		router.GET("/user/:id", userController.GetUserById)
		router.GET("/user", userController.GetUsers)

		// Login Mapping
		router.POST("/login", loginController.Login)
	*/
	router.POST("/Singup", userController.Singup)
	router.POST("/login", userController.Login)
	router.POST("/logout", userController.Logout)
	router.GET("/validate", auth.RequireAuth, userController.Validate)
	// User Mapping
	router.GET("/user/id/:id", auth.RequireAuth, userController.GetUserById)
	router.GET("/user/username/:name", userController.GetUSerByusername)
	// hotel Mapping
	router.GET("/hotel/id/:id", hotelController.GetHotelById)

	// bookings Mapping
	router.GET("/bookings/mybookings/:id", bookingController.GetmyBookings)
	router.GET("/bookings/bookings/", auth.RequireAdmin, bookingController.GetBookings)
	//router.GET("/bookings/id/:id", bookingController.GetBookingsById)
	router.POST("/reserve/:id", auth.RequireAuth, bookingController.Reserve)
	// availablehotels Mapping
	router.GET("/hotels/", hotelController.GetHotels)
	router.POST("/availablehotels/", AvailableordersController.GetAvailableHotels)
}
