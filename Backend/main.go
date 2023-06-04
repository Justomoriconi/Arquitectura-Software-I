package main

import (
	"Backend/Database"
	"Backend/app"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// prueba
func main() {

	Database.StartDbEngine()
	app.StartRoute()
}
