package controllers

import (
	"encoding/json"
	"net/http"

	"../models"
)

func GetUsersEndpoint(w http.ResponseWriter, req *http.Request) {
	users := models.GetUsers()
	json.NewEncoder(w).Encode(users)
}
