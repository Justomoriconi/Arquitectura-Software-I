package Model

type User struct {
	UserID   int    `gorm:"primary_key;autoIncrement:true;unique"`
	Name     string `gorm:"type:varchar(255);not null"`
	LastName string `gorm:"type:varchar(255);not null"`
	UserName string `gorm:"type:varchar(255)"`
	Email    string `gorm:"type:varchar(255);not null;unique"`
	Pwd      string `gorm:"type:varchar(255);not null"`
	Admin    bool   `gorm:"type:boolean;default:0"`
}
type Users []User
