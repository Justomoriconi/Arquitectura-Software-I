package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

import "Backend/Database"

func main() {
	Database.StartDbEngine()
}

