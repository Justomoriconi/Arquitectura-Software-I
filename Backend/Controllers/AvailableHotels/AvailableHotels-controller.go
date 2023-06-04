package Controllers

import (
	service "Backend/Services/AvailableHotels"

	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAvailableHotels(c *gin.Context) {
	// Retrieve check-in and check-out dates from the query parameters
	checkin := c.Query("checkin")
	checkout := c.Query("checkout")

	availableHotels, err := service.HotelServices.GetAvailableHotels(checkin, checkout)

	if err != nil {
		c.JSON(http.StatusBadRequest, availableHotels)
		return
	}

	// Return the available hotels in the response
	c.JSON(http.StatusOK, availableHotels)
}