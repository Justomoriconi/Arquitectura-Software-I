package Controllers

import (
	service "Backend/Services/Bookings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func GetBookingsById(c *gin.Context) {
	log.Debug("Bookings id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	bookingdomain, err := service.BookingService.GetBookinglById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, bookingdomain)
		return
	}

	c.JSON(http.StatusOK, bookingdomain)
}

func GetmyBookings(c *gin.Context) {
	log.Debug("Bookings from id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	bookingdomain, err := service.BookingService.GetmyBookings(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, bookingdomain)
		return
	}

	c.JSON(http.StatusOK, bookingdomain)
}

/*
func Reserve(c *gin.Context) {

	user := c.Param("user")
	//we are using auth as a midleware giving user so we now it exist

	bookingdomain, err := service.BookingService.GetmyBookings(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, bookingdomain)
		return
	}

	c.JSON(http.StatusOK, bookingdomain)
}
*/
