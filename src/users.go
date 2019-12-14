package main

import (
	"database/sql"
	"log"
)

func GetUserById(userid int, db *sql.DB) User {

	var password string // we don't do anything with this

	var user User

	e := db.QueryRow("SELECT * FROM users WHERE id=$1;",
		userid).Scan(&user.Id, &user.Age, &user.Firstname, &user.Lastname, &user.Email,
			&user.Username, &user.Public, &user.Joindate, &user.Active, &password, &user.Gender)
	if e != nil {
		log.Println(Warn("Unable to fetch user from database."))
		return User{}
	}

	return user

}