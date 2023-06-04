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

func GetHotelByname(c *gin.Context) {
	log.Debug("Hotel name to load: " + c.Param("name"))

	name := c.Param("name")
	hoteldomain, err := service.HotelService.GetHotelByname(name)

	if err != nil {
		c.JSON(http.StatusBadRequest, hoteldomain)
		return
	}

	c.JSON(http.StatusOK, hoteldomain)
}
