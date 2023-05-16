package usersdto

type CreateUserRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"fullName" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Phone    string `json:"phone" form:"phone"`
	Location string `json:"location" form:"location"`
	Image    string `json:"image" form:"image"`
	Role     string `json:"role" form:"role"`
}

type UpdateUserRequest struct {
	Name     string `json:"fullName" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Gender   string `json:"gender" gorm:"type: varchar(255)" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Location string `json:"location" form:"location" validate:"required"`
	Image    string `json:"image" form:"image" validate:"required"`
}
