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
	//admin email:"admin@gmail.com"
	//pwd :secret

	if err != nil {
		db.Create(&userModel.User{Admin: true, Name: "Admin", LastName: "Admin", UserName: "Admin", Email: "admin@gmail.com", Pwd: "$2a$10$SNYWKch/a88Z9Ql1OmJYzeUKt3CvmVsXBD2H1HEVT4gRRxl7iDp1K"})

	}
	//Inserting Hotels

	err = db.First(&hotelModel.Hotel{}).Error

	if err != nil {
		db.Create(&hotelModel.Hotel{Name: "Gardi Hotel & Suites", Rooms: 1, Description: " El Gardi Hotel & Suites se encuentra en Buenos Aires, a 400 metros del teatro Colón, y ofrece alojamiento con bar, estacionamiento privado y jardín."})
		db.Create(&hotelModel.Hotel{Name: "Ayres Recoleta Uriburu", Rooms: 2, Description: "El Ayres Recoleta Uriburu se encuentra en Buenos Aires, a 400 metros del cementerio de la Recoleta, y ofrece alojamiento con salón compartido, wifi gratis, centro de negocios y servicio de conserjería. Se encuentra a menos de 1 km del Museo Nacional de Bellas Artes y ofrece un mostrador de información turística."})
		db.Create(&hotelModel.Hotel{Name: "Clayton Buenos Aires", Rooms: 3, Description: "El Clayton Buenos Aires goza de una buena ubicación en el barrio de Palermo de Buenos Aires, a menos de 1 km de Bosques de Palermo, a 8 minutos a pie de los lagos de Palermo y a menos de 1 km del parque El Rosedal. El alojamiento se encuentra a 1,5 km del jardín japonés de Buenos Aires y a 2,7 km del Museo de Arte Latinoamericano de Buenos Aires MALBA y de la plaza Serrano. Hay recepción 24 horas, salón compartido y servicio de cambio de divisa."})
		db.Create(&hotelModel.Hotel{Name: "Trendy & Luxury", Rooms: 5, Description: "El Trendy & Luxury se encuentra en Buenos Aires, en el barrio de Palermo, a 500 metros de la plaza Serrano y a 2,3 km del parque ecológico de Buenos Aires, y ofrece jardín. El departamento tiene terraza."})
		db.Create(&hotelModel.Hotel{Name: "Hotel Regis", Rooms: 6, Description: "El Hotel Regis está bien situado en el centro de Buenos Aires, a 800 metros del Obelisco de Buenos Aires, a menos de 1 km del teatro Colón y a 13 minutos a pie de la basílica del Santísimo Sacramento."})

	}

	//Inserting bookings
	err = db.First(&bookingModel.Booking{}).Error

	if err != nil {
		db.Create(&bookingModel.Booking{UserID: 1, HotelID: 1, Checkin: "2023-05-20", Checkout: "2023-05-25"})
		db.Create(&bookingModel.Booking{UserID: 1, HotelID: 2, Checkin: "2023-05-21", Checkout: "2023-05-25"})
		db.Create(&bookingModel.Booking{UserID: 2, HotelID: 2, Checkin: "2023-06-21", Checkout: "2023-07-25"})
		db.Create(&bookingModel.Booking{UserID: 2, HotelID: 3, Checkin: "2020-06-21", Checkout: "2023-07-25"})

	}

	log.Info("Data inserted")
}
