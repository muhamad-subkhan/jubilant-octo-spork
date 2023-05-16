package repositories

import (
	"BE-foodways/models"

	"gorm.io/gorm"
)

type UserRepositories interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(user models.User) (models.User, error)
}

func RepositoriesUser(db *gorm.DB) *repositories {
	return &repositories{db}

}

func (r *repositories) FindUsers() ([]models.User, error) {
	var users []models.User
	// err := r.db.Raw("SELECT * FROM users").Scan(&users).Error
	err := r.db.Find(&users).Error

	return users, err
}

func (r *repositories) GetUser(ID int) (models.User, error) {
	var user models.User
	// err := r.db.Raw("SELECT * FROM users WHERE id=?", ID).Scan(&user).Error
	err := r.db.First(&user, ID).Error

	return user, err
}

func (r *repositories) CreateUser(user models.User) (models.User, error) {
	// err := r.db.Exec("INSERT INTO users(name, email, phone, location, image, role) VALUES (?,?,?,?,?,?)", user.Name, user.Email, user.Phone, user.Location, user.Image, user.Role).Error
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repositories) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repositories) DeleteUser(user models.User) (models.User, error) {
	err := r.db.Delete(&user).Error

	return user, err
}
