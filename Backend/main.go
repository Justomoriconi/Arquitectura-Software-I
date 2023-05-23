package main

import (
	"Backend/Database"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	Database.StartDbEngine()
}
