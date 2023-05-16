package repositories

import (
	"BE-foodways/models"

	"gorm.io/gorm"
)

type AuthRepositories interface {
	Register(user models.User) (models.User, error)
	Login(email string) (models.User, error)
	GetUser(ID int) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *repositories {
	return &repositories{db}
}

func (r *repositories) Register(abc models.User) (models.User, error) {
	err := r.db.Create(&abc).Error
	return abc, err
}

func (r *repositories) Login(email string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "email=?", email).Error

	return user, err
}

func (r *repositories) Getuser(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error

	return user, err
}
