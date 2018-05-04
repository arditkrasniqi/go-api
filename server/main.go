package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"./controllers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", controllers.GetUsersEndpoint).Methods("GET")
	router.HandleFunc("/delete-user/{id}", controllers.DeleteUserEndpoint).Methods("DELETE")
	router.HandleFunc("/update-user/{id}", controllers.UpdateUserEndpoint).Methods("PATCH")
	router.HandleFunc("/new-user", controllers.NewUserEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":9000", router))
}
