package Clients

import (
	model "Backend/Model/User"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetUserById(id int) (model.User, error) { //standard get by id
	var user model.User

	err := Db.Where("user_id = ?", id).First(&user).Error

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
func GetUserByEmail(Email string) (model.User, error) { //looking for username to log in
	var user model.User

	err := Db.Where("email = ?", Email).First(&user).Error

	if err != nil { //is empty
		log.Println(err)
		return user, err
	}
	log.Debug("User: ", user)

	return user, nil
}
func PutUser(User *model.User) error {
	err := Db.Create(User).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}
