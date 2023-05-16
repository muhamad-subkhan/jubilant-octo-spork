package models

type User struct {
	ID       int    `json:"id" gorm:"primary_key:auto_increment"`
	Name     string `json:"fullName" gorm:"type: varchar (255)"`
	Password string `json:"password" gorm:"type: varchar (255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Gender   string `json:"gender" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" gorm:"type: varchar(25)"`
	Location string `json:"location" gorm:"type: varchar(255)"`
	Image    string `json:"image" gorm:"type: varchar(255)"`
	Role     string `json:"role" gorm:"type: varchar(10)"`
}

type UsersProfileResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"fullName"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Location string `json:"location"`
}

type UserOrderResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"fullName"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Location string `json:"location"`
}
func (UsersProfileResponse) TableName() string {
	return "users"
}

func (UserOrderResponse) TableName() string {
	return "users"
}