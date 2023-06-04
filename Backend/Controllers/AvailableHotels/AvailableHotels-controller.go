package Controllers

/*
import (
	service "Backend/Services/AvailableHotels"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func GetAvailableHotels(c *gin.Context) {
	log.Debug("Hotel id to load: " + c.JSON("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	hoteldomain, err := service.HotelService.GetHotelById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, hoteldomain)
		return
	}

	c.JSON(http.StatusOK, hoteldomain)
}*/
