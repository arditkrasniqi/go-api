package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"

	"../models"
	"log"
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

func UpdateUserEndpoint(w http.ResponseWriter, req *http.Request){
	decoder := json.NewDecoder(req.Body)
	var user *models.User
	err := decoder.Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	params := mux.Vars(req)
	users := models.UpdateUser(params["id"], *user)
	json.NewEncoder(w).Encode(users)
}

func NewUserEndpoint(w http.ResponseWriter, req *http.Request){
	decoder := json.NewDecoder(req.Body)
	var user *models.User
	err := decoder.Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	users := models.CreateUser(*user)
	json.NewEncoder(w).Encode(users)
}