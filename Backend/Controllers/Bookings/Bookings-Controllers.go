package Controllers

import (
	userModel "Backend/Model/User"
	service "Backend/Services/Bookings"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
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
	}
*/
func GetmyBookings(c *gin.Context) {
	log.Debug("loading my bookings: ")
	//id, _ := strconv.Atoi(c.Param("id"))
	User, err := c.Get("User")
	if err != true {
		c.JSON(http.StatusBadRequest, gin.H{"not user": err})
	}
	c.JSON(http.StatusUnauthorized, gin.H{"User": User.(userModel.User).UserID})
	/*
		bookingdomain, err := service.BookingService.GetmyBookings(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, bookingdomain)
			return
		}*/

	c.JSON(http.StatusOK, gin.H{"hola": err})
}

/*
func Reserve(c *gin.Context) {
	log.Debug("hotel to reserve: " + c.Param("hotelid"))
	user, err := c.Get("user")
	if err! = nil {
		c.JSON(http.StatusBadRequest, gin.Context{})
		return
	}

	id := user.(model.User).UserID
	checkin := c.Query("checkin")
	checkout := c.Query("checkout")
	hotelid, _ := strconv.Atoi(c.Param("hotelid"))
	bookingdomain, err := service.BookingService.Reserve(id, hotelid, checkin, checkout)

	if err != nil {
		c.JSON(http.StatusBadRequest, bookingdomain)
		return
	}

	c.JSON(http.StatusOK, bookingdomain)
}
*/
