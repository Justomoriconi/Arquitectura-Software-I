package Controllers

import (
	service "Backend/Services/AvailableHotels"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAvailableHotels(c *gin.Context) {
	// Retrieve check-in and check-out dates from the Body parameters
	
	var Body struct {
		Checkin  string
		Checkout string
	}

	if c.Bind(&Body) != nil {
		c.JSON(http.StatusBadRequest, Body)
		return
	}

	availableHotels, err := service.HotelServices.GetAvailableHotels(Body.Checkin, Body.Checkout)

	if err != nil {
		c.JSON(http.StatusBadRequest, availableHotels)
		return
	}

	// Return the available hotels in the response
	c.JSON(http.StatusOK, availableHotels)
}
