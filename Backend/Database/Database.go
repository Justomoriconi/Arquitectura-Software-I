package Database

import (
	data "Backend/Database/data"
	bookingModel "Backend/Model/Booking"
	hotelModel "Backend/Model/Hotel"
	userModel "Backend/Model/User"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "proyect" //variable de entorno para nombre de la base de datos
	DBUser := "user"    //variable de entorno para el usuario de la base de datos
	//DBPass := ""
	DBPass := "secret" //variable de entorno para la pass de la base de datos
	DBHost := "localhost"
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	db.LogMode(true)

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	/*	// We need to add all CLients that we build
		userClient.Db = db
		productClient.Db = db
		addressClient.Db = db
		orderClient.Db = db
		searchClient.Db = db
		homeClient.Db = db*/

}

func StartDbEngine() {

	// We need to migrate all classes model.
	db.AutoMigrate(&userModel.User{}, &hotelModel.Hotel{}, &bookingModel.Booking{})

	log.Info("Finishing Migration Database Tables")
	data.InsertData(db)
}
