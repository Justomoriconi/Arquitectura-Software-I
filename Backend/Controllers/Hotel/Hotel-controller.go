package Controllers

import (
	service "Backend/Services/Hotel"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func GetHotelById(c *gin.Context) {
	log.Debug("Hotel id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	hoteldomain, err := service.HotelService.GetHotelById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, hoteldomain)
		return
	}

	c.JSON(http.StatusOK, hoteldomain)
}

func GetHotels(c *gin.Context) {
	log.Debug("loading hotels: ")

	hoteldomain, err := service.HotelService.GetHotels()

	if err != nil {
		c.JSON(http.StatusBadRequest, hoteldomain)
		return
	}

	c.JSON(http.StatusOK, hoteldomain)
}
