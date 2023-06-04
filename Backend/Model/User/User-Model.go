package Model

type User struct {
	UserID   int    `gorm:"primary_key;autoIncrement:true;unique"`
	Name     string `gorm:"type:varchar(255);not null"`
	LastName string `gorm:"type:varchar(255);not null"`
	UserName string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null"`
	Pwd      string `gorm:"type:varchar(255);not null"`
}
type Users []User
