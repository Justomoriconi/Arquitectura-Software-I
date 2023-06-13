package Controllers

import (
	service "Backend/Services/Bookings"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func GetBookings(c *gin.Context) {
	log.Debug("loading bookings: ")

	hoteldomain, err := service.BookingService.GetBookings()

	if err != nil {
		c.JSON(http.StatusBadRequest, hoteldomain)
		return
	}

	c.JSON(http.StatusOK, hoteldomain)
}

/*
func GetmyBookings(c *gin.Context) {
	log.Debug("loading my bookings: ")
	user, _ := c.Get("user")
	id := user.(model.User).UserID

	bookingdomain, err := service.BookingService.GetmyBookings(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, bookingdomain)
		return
	}

	c.JSON(http.StatusOK, bookingdomain)
}*/

func GetmyBookings(c *gin.Context) {
	log.Debug("loading my bookings: ")
	id, _ := strconv.Atoi(c.Param("id"))

	bookingdomain, err := service.BookingService.GetmyBookings(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, bookingdomain)
		return
	}

	c.JSON(http.StatusOK, bookingdomain)
}

func Reserve(c *gin.Context) {
	var Body struct {
		Userid   string
		Hotelid  string
		Checkin  string
		Checkout string
	}
	if c.Bind(&Body) != nil {
		c.JSON(http.StatusBadRequest, Body)
		return
	}
	id, _ := strconv.Atoi(Body.Userid)
	hotelid, _ := strconv.Atoi(Body.Hotelid)
	booking, err := service.BookingService.Reserve(id, hotelid, Body.Checkin, Body.Checkout)

	if err != nil {
		c.JSON(http.StatusBadRequest, booking)
		return
	}

	c.JSON(http.StatusOK, booking)
}
