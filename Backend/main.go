package main

import (
	hotelcontroler "Backend/Controllers/Hotel"
	"Backend/Database"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// prueba
func main() {
	router := gin.Default()
	Database.StartDbEngine()
	router.GET("/hotel/id/:id", hotelcontroler.GetHotelById)
	router.GET("/hotel/name/:name", hotelcontroler.GetHotelByname)
	router.Run()
}
