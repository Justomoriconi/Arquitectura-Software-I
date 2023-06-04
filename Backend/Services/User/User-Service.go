package User

/*
import (
	client "Backend/Clients/User"
	domain "Backend/Domain/User"
)

type userService struct{}

type UserServiceInterface interface {
	GetUserById(int) (domain.User, error)
	GetUSerByusername(string) (domain.Hotel, error)
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
	Userdomain.Pwd=
	return Userdomain, nil
}

func (s *hotelService) GetHotelByname(name string) (domain.Hotel, error) {

	hotel, err := client.GetHotelByName(name)
	var HotelDomain domain.Hotel

	if err != nil {
		return HotelDomain, err
	}
	if hotel.HotelID == 0 {
		return HotelDomain, nil
	}
	HotelDomain.Name = hotel.Name
	HotelDomain.ID = hotel.HotelID
	HotelDomain.Rooms = hotel.Rooms

	return HotelDomain, nil
}
*/
