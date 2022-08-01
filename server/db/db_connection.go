package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	DB_DSN = "postgres://localhost:5432/fruits?sslmode=disable"
)

var DB *sql.DB

func Db_connection() {
	var err error
	DB, err = sql.Open("postgres", DB_DSN)

	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	} else {
		fmt.Println("connected")
	}
	// defer DB.Close()

}
