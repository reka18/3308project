package main

import (
	"database/sql"
	"log"
)

func GetUserId(username string, db *sql.DB) int {

	var id int
	e := db.QueryRow("SELECT id FROM users WHERE username=$1;", username).Scan(&id)
	if e != nil {
		log.Println(Warn("Failed to get id for referenced user."))
	}
	return id

}