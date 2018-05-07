package router

import (
	"../../controllers"
	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router) {
	router.HandleFunc("/users", controllers.GetUsersEndpoint).Methods("GET")
	router.HandleFunc("/delete-user/{id}", controllers.DeleteUserEndpoint).Methods("DELETE")
	router.HandleFunc("/update-user/{id}", controllers.UpdateUserEndpoint).Methods("PATCH")
	router.HandleFunc("/new-user", controllers.NewUserEndpoint).Methods("POST")
}