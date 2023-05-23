package Data

import (
	bookingModel "Backend/Model/Booking"
	hotelModel "Backend/Model/Hotel"
	userModel "Backend/Model/User"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

func InsertData(db *gorm.DB) {
	// Insert data
	log.Info("Inserting data...")

	//Inserting users
	err := db.First(&userModel.User{}).Error

	if err != nil {
		db.Create(&userModel.User{UserID: 1, Name: "Augusto", LastName: "Bruno", UserName: "augustob", Email: "abcdefg@gmail.com", Pwd: "hola123"})
	}
	//Inserting Hotels

	err = db.First(&hotelModel.Hotel{}).Error

	if err != nil {
		db.Create(&hotelModel.Hotel{HotelID: 1, Name: "hotel1", Rooms: 4})
		db.Create(&hotelModel.Hotel{HotelID: 2, Name: "hotel2", Rooms: 3})
		db.Create(&hotelModel.Hotel{HotelID: 3, Name: "hotel3", Rooms: 2})

	}

	//Inserting bookings
	err = db.First(&bookingModel.Booking{}).Error

	if err != nil {
		db.Create(&bookingModel.Booking{UserID: 1, HotelID: 1})
		db.Create(&bookingModel.Booking{UserID: 1, HotelID: 2})

	}

	//manage errors...

	log.Info("Data inserted")
}
