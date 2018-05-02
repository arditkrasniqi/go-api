package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"

	"../models"
)

func GetUsersEndpoint(w http.ResponseWriter, req *http.Request) {
	users := models.GetUsers()
	json.NewEncoder(w).Encode(users)
}

func DeleteUserEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	users := models.DeleteUser(params["id"])
	json.NewEncoder(w).Encode(users)
}
