package routes

import (
	"BE-foodways/handlers"
	"BE-foodways/pkg/middleware"
	"BE-foodways/pkg/mysql"
	"BE-foodways/repositories"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {

	ProductsRepository := repositories.RepositoriesProduct(mysql.DB)
	h := handlers.HandlerProduct(ProductsRepository)
  
	r.HandleFunc("/products", h.FindProducts).Methods("GET") //get all
	r.HandleFunc("/products/{id}", middleware.Auth(h.GetProduct)).Methods("GET") //get by id
	r.HandleFunc("/product/{id}", middleware.Auth(h.ProductDetail)).Methods("GET") //get detail by id
	r.HandleFunc("/product", middleware.Auth(middleware.UploadFile(h.CreateProduct))).Methods("POST") //Create
	r.HandleFunc("/products/{id}", middleware.Auth(h.UpdateProduct)).Methods("PATCH") // Update Product
	r.HandleFunc("/product/{id}", middleware.Auth(h.DeleteProduct)).Methods("DELETE") // Delete Product


}