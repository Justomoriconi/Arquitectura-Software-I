package Clients

import (
	model "Backend/Model/User"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetUserById(id int) (model.User, error) { //standard get by id
	var user model.User

	err := Db.Where("UserID = ?", id).First(&user).Error

	if err != nil {
		log.Println(err)
		return user, err
	}

	log.Debug("User: ", user)

	return user, nil
}

func GetUserByUsername(username string) (model.User, error) { //looking for username to log in
	var user model.User

	err := Db.Where("user_id = ?", username).First(&user).Error

	if err != nil { //is empty
		log.Println(err)
		return user, err
	}
	log.Debug("User: ", user)

	return user, nil
}

func GetUserByEmail(email string) (model.User, error) { //looking for username to log in
	var user model.User

	err := Db.Where("user_email = ?", email).First(&user).Error

	if err != nil { //is empty
		log.Println(err)
		return user, err
	}
	log.Debug("User: ", user)

	return user, nil
}
