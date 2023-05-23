package Models

type User struct {
	UserID   int    `json:"ID"`
	Name     string `json:"Name"`
	LastName string `gorm:"Lastname"`
	UserName string `gorm:"Username"`
	Email    string `gorm:"Email"`
	Pwd      string `gorm:"Pasword"`
}
