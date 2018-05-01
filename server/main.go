package main

import (

	// "encoding/json"

	"log"
	"net/http"

	"./controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", controllers.GetUsersEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000", router))
}
