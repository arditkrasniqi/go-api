package main

import (
	"database/sql"
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
    "time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

type User struct{
    ID int `json:"id,omitempty"`
    Username string `json:"username,omitempty"`
    Password string `json:"password,omitempty"`
}

type Post struct{
    ID int `json:"id,omitempty"`
    PostedBy int `json:"postedby,omitempty"`
    Title string `json:"title,omitempty"`
    Content string `json:"content,omitempty"`
    PostedAt time.Time `json:"postedat,omitempty"`
}

func Test(w http.ResponseWriter, req *http.Request){
    fmt.Fprint(w,"Testing")
}

func main(){
    db, _ = sql.Open("mysql", "root:password@tcp(localhost:3306)/portal")
	dberror := db.Ping()
	if dberror != nil {
		log.Fatal(dberror)
	}
    router := mux.NewRouter()
    router.HandleFunc("/test", Test).Methods("GET")
    log.Fatal(http.ListenAndServe(":9000", router))
}
