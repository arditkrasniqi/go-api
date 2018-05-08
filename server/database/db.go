package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Open() *sql.DB {
	if db != nil {
		return db
	} else {
		db, _ = sql.Open("mysql", "root:@tcp(localhost:3306)/portal")
		dberror := db.Ping()
		if dberror != nil {
			log.Fatal(dberror)
		}
		return db
	}
}

func Close() {
	if db != nil {
		db.Close()
	}
}
