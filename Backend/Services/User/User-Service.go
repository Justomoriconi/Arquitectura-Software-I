package User

import (
	client "Backend/Clients/User"
	domain "Backend/Domain/User"
	model "Backend/Model/User"
)

type userService struct{}

type UserServiceInterface interface {
	GetUserById(int) (domain.User, error)
	GetUSerByEmail(email string) (domain.User, error)
	GetUSerByusername(string) (domain.User, error)
	Singup(Name string, LastName string, UserName string, Email string, Pwd string) error
}

var (
	UserService UserServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id int) (domain.User, error) {

	user, err := client.GetUserById(id)
	var Userdomain domain.User

	if err != nil {
		return Userdomain, err
	}
	if user.UserID == 0 {
		return Userdomain, nil
	}
	Userdomain.Name = user.Name
	Userdomain.UserID = user.UserID
	Userdomain.UserName = user.UserName

	return Userdomain, nil
}

func (s *userService) GetUSerByusername(username string) (domain.User, error) {

	user, err := client.GetUserByUsername(username)
	var Userdomain domain.User

	if err != nil {
		return Userdomain, err
	}
	if user.UserID == 0 {
		return Userdomain, nil
	}
	Userdomain.Name = user.Name
	Userdomain.UserID = user.UserID
	Userdomain.UserName = user.UserName

	return Userdomain, nil
}

func (s *userService) GetUSerByEmail(email string) (domain.User, error) {

	user, err := client.GetUserByEmail(email)
	var Userdomain domain.User

	if err != nil {
		return Userdomain, err
	}
	if user.UserID == 0 {
		return Userdomain, nil
	}
	Userdomain.Name = user.Name
	Userdomain.Name = user.LastName
	Userdomain.UserID = user.UserID
	Userdomain.UserName = user.UserName
	Userdomain.Pwd = user.Pwd
	return Userdomain, nil
}

func (s *userService) Singup(Name string, LastName string, UserName string, Email string, Pwd string) error {
	var Usermodel model.User
	Usermodel.Pwd = Pwd
	Usermodel.Name = Name
	Usermodel.LastName = LastName
	Usermodel.Email = Email
	Usermodel.UserName = UserName

	err := client.PutUser(&Usermodel)

	return err

}
