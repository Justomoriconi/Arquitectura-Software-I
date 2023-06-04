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
		db.Create(&userModel.User{Name: "Augusto", LastName: "Bruno", UserName: "augustob", Email: "abcdefg@gmail.com", Pwd: "hola123"})
		db.Create(&userModel.User{Name: "Justo", LastName: "Moriconi", UserName: "justom", Email: "abcdefg@gmail.com", Pwd: "hola123"})

	}
	//Inserting Hotels

	err = db.First(&hotelModel.Hotel{}).Error

	if err != nil {
		db.Create(&hotelModel.Hotel{Name: "hotel1", Rooms: 1})
		db.Create(&hotelModel.Hotel{Name: "hotel2", Rooms: 2})
		db.Create(&hotelModel.Hotel{Name: "hotel3", Rooms: 3})
		db.Create(&hotelModel.Hotel{Name: "hotel4", Rooms: 5})
		db.Create(&hotelModel.Hotel{Name: "hotel5", Rooms: 6})
		db.Create(&hotelModel.Hotel{Name: "hotel6", Rooms: 7})
		db.Create(&hotelModel.Hotel{Name: "hotel7", Rooms: 4})
		db.Create(&hotelModel.Hotel{Name: "hotel8", Rooms: 1})
		db.Create(&hotelModel.Hotel{Name: "hotel9", Rooms: 3})

	}

	//Inserting bookings
	err = db.First(&bookingModel.Booking{}).Error

	if err != nil {
		db.Create(&bookingModel.Booking{UserID: 1, HotelID: 1, Checkin: "2023-05-20", Checkout: "2023-05-25"})
		db.Create(&bookingModel.Booking{UserID: 1, HotelID: 2, Checkin: "2023-05-21", Checkout: "2023-05-25"})
		db.Create(&bookingModel.Booking{UserID: 2, HotelID: 2, Checkin: "2023-06-21", Checkout: "2023-07-25"})
		db.Create(&bookingModel.Booking{UserID: 2, HotelID: 8, Checkin: "2020-06-21", Checkout: "2023-07-25"})

	}

	log.Info("Data inserted")
}
