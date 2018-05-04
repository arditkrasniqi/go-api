package models

import (
	"log"

	"../database"
)

type User struct {
	ID       string `json:"id,omitempty"`
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
	return users
}

func DeleteUser(id string) []User {
	_, err := db.Query("delete from users where id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	return GetUsers()
}

func UpdateUser(id string, user User) []User {
	_, err := db.Query("update users set username = ?, password = ? where id = ?", user.Username, user.Password, id)
	if err != nil {
		log.Fatal(err)
	}
	return GetUsers()
}

func CreateUser(user User) []User{
	_, err := db.Query("insert into users(`username`,`password`) values(?,?)", user.Username, user.Password)
	if err != nil{
		log.Fatal(err)
	}
	return GetUsers()
}
