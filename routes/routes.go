package routes

import "github.com/gorilla/mux"

func Routes(r *mux.Router){
	UserRoutes(r)
	AuthRoutes(r)
	ProductRoutes(r)
}