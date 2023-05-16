package routes

import (
	"BE-foodways/handlers"
	"BE-foodways/pkg/middleware"
	"BE-foodways/pkg/mysql"
	"BE-foodways/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {

	UserRepository := repositories.RepositoriesUser(mysql.DB)
	h := handlers.HandlerUser(UserRepository)

	r.HandleFunc("/users", h.FindUsers).Methods("GET")         //Get All
	r.HandleFunc("/users/{id}", h.GetUser).Methods("GET")      //get by id
	r.HandleFunc("/user", h.CreateUser).Methods("POST")        // Create
	r.HandleFunc("/user/{id}",middleware.UploadFile( h.UpdateUser)).Methods("PATCH")  // Update
	r.HandleFunc("/user/{id}", h.DeleteUser).Methods("DELETE") // Delete

}
