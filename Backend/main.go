package main

import (
	"Backend/Database"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// prueba
func main() {
	Database.StartDbEngine()
}
