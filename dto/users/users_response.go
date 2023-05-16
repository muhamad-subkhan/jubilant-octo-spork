package usersdto

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"fullName" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	// Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
	Phone    string `json:"phone" form:"phone"`
	Gender   string `gorm:"type: varchar(255)" json:"gender" validate:"required"`
	Location string `json:"location" form:"location"`
	Image    string `json:"image" form:"image"`
	Role     string `json:"role" form:"role" validate:"required"`
}

// type GetUser struct{
// 	ID       int    `json:"id"`
// 	Name     string `json:"fullName" form:"name" validate:"required"`
// 	Email    string `json:"email" form:"email" validate:"required"`
// 	Phone    string `json:"phone" form:"phone"`
// 	Gender   string `gorm:"type: varchar(255)" json:"gender" validate:"required"`
// 	Location string `json:"location" form:"location"`
// 	Image    string `json:"image" form:"image"`
// 	Role     string `json:"role" form:"role" validate:"required"`
// }

type UpdateRespone struct {
	ID       int    `json:"id"`
	Name     string `json:"fullName" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Phone    string `json:"phone" form:"phone"`
	Gender   string `gorm:"type: varchar(255)" json:"gender" validate:"required"`
	Location string `json:"location" form:"location"`
	Image    string `json:"image" form:"image"`
	Role     string `json:"role" form:"role" validate:"required"`
}