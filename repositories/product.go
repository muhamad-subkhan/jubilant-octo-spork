package repositories

import (
	"BE-foodways/models"

	"gorm.io/gorm"
)

type ProductRepositories interface {
	FindProducts() ([]models.Product, error)
	GetProduct(ID int) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	ProductDetail(ID int) (models.Product, error)
	UpdateProduct(product models.Product) (models.Product, error)
	DeleteProduct(product models.Product) (models.Product, error)
}

func RepositoriesProduct(db *gorm.DB) *repositories {
	return &repositories{db}
}

func (r *repositories) FindProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("User").Find(&products).Error

	return products, err
}

func (r *repositories) GetProduct(ID int) (models.Product, error) {
	var product models.Product
	// not yet using category relation, cause this step doesnt Belong to Many
	err := r.db.Preload("User").First(&product, ID).Error

	return product, err
}

func (r *repositories) CreateProduct(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error

	return product, err
}

func (r *repositories) ProductDetail(ID int) (models.Product, error) {
	var product models.Product
	// not yet using category relation, cause this step doesnt Belong to Many
	err := r.db.Preload("User").First(&product, ID).Error

	return product, err
}

func (r *repositories) UpdateProduct(product models.Product) (models.Product, error) {
	err := r.db.Save(&product).Error

	return product, err
}

func (r *repositories) DeleteProduct(product models.Product) (models.Product, error) {
	err := r.db.Delete(&product).Error

	return product, err
}