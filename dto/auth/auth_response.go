package authdto

type LoginResponse struct {
	ID    int    `gorm:"type: int" json:"id"`
	Name  string `gorm:"type: varchar(255)" json:"name"`
	Email string `gorm:"type: varchar(255)" json:"email"`
	// Password string `gorm:"type: varchar(255)" json:"password"`
	Token string `gorm:"type: varchar(255)" json:"token"`
	Role  string `gorm:"type: varchar(10)" json:"role"`
}

type CheckAuthResponse struct {
	Id    int    `gorm:"type: int" json:"id"`
	Name  string `gorm:"type: varchar(255)" json:"name"`
	Email string `gorm:"type: varchar(255)" json:"email"`
	Role  string `gorm:"type: varchar(10)"  json:"role"`
}