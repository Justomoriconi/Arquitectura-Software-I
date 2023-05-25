package Clients

import (
	model "Backend/Model/Hotel"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetHotelById(id int) (model.Hotel, error) { //standard get by id
	var hotel model.Hotel

	err := Db.Where("hotel_id = ?", id).First(&hotel).Error

	if err != nil {
		log.Println(err)
		return hotel, err
	}

	log.Debug("Hotel: ", hotel)

	return hotel, nil
}
