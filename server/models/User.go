package models

import (
	"log"

	"../database"
)

type User struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

var db = database.Open()

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) GetPassword() string {
	return u.Password
}

func GetUsers() []User {
	var users []User
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	model := User{}
	for rows.Next() {
		if err := rows.Scan(&model.ID, &model.Username, &model.Password); err != nil {
			log.Fatal(err)
		}
		users = append(users, model)
	}
	defer rows.Close()
	defer db.Close()
	return users
}
