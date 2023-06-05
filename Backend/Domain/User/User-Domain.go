package Domain

type User struct {
	UserID   int    `json:"ID"`
	Name     string `json:"Name"`
	LastName string `json:"Lastname"`
	UserName string `json:"Username"`
	Email    string `json:"Email"`
	Pwd      string `json:"Pasword"`
}
type Users []User
