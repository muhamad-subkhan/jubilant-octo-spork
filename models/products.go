package models

type Product struct {
	ID     int                  `json:"id" gorm:"primary_key:auto_increment"`
	Title  string               `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price  int                  `json:"price" form:"price" gorm:"type: int"`
	Image  string               `json:"image" form:"image" gorm:"type: varchar(255)"`
	UserID int                  `json:"user_id" form:"user_id"`
	User   UsersProfileResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ProductResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
	Image string `json:"image"`
}
type ProductDelete struct {
	ID int `json:"id"`
}

type OrderResponse struct {
	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
	Title string `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price int    `json:"price" form:"price" gorm:"type: int"`
	Image string `json:"image" form:"image" gorm:"type: varchar(255)"`
}

type ProductDetail struct {
	ID     int                  `json:"id"`
	Title  string               `json:"title"`
	Price  int                  `json:"price"`
	Image  string               `json:"image"`
	UserID int                  `json:"-"`
	User   UsersProfileResponse `json:"user"`
}

// type ProductUserResponse struct {
// 	ID     int    `json:"id"`
// 	Name   string `json:"name"`
// 	Desc   string `json:"desc"`
// 	Price  int    `json:"price"`
// 	Image  string `json:"image"`
// 	Qty    int    `json:"qty"`
// 	UserID int    `json:"-"`
// }

func (ProductResponse) TableName() string {
	return "products"
}
func (ProductDelete) TableName() string {
	return "products"
}

func (OrderResponse) TableName() string {
	return "products"
}

// func (ProductUserResponse) TableName() string {
// 	return "products"
// }