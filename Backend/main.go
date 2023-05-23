package main

<<<<<<< Updated upstream
import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
func init() {

	DBName := "proyect"      //variable de entorno para nombre de la base de datos
	DBUser := "proyect_user" //variable de entorno para el usuario de la base de datos
	DBPass := "secret"       //variable de entorno para la pass de la base de datos
	DBHost := "db"

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	db.LogMode(true)

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}
*/
func main() {
	_, err = db.Exec("CREATE DATABASE mydatabase")
	if err != nil {
		panic(err)
	}

	fmt.Println("Database created successfully!")
}
}
=======
import "Backend/Database"

func main() {
	Database.StartDbEngine()
}
>>>>>>> Stashed changes
